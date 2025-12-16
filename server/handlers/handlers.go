package handlers

import (
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/pkg/test"
)

// TestExecutionRequest represents the request body for test execution
// Now accepts the suite configuration directly instead of a file path
type TestExecutionRequest struct {
	Suite              TestSuiteRequest `json:"suite" binding:"required"`
	ApiKey             string           `json:"api_key,omitempty"`             // Optional token to override global.VoiceflowAPIKey
	VoiceflowSubdomain string           `json:"voiceflow_subdomain,omitempty"` // Optional subdomain to override global.VoiceflowSubdomain
} // @name TestExecutionRequest

// TestSuiteRequest represents a test suite configuration for HTTP requests
type TestSuiteRequest struct {
	Name              string        `json:"name" binding:"required" example:"Example Suite"`
	Description       string        `json:"description" example:"Suite used as an example"`
	EnvironmentName   string        `json:"environment_name" binding:"required" example:"production"`
	NewSessionPerTest bool          `json:"new_session_per_test" example:"false"`
	Tests             []TestRequest `json:"tests" binding:"required,dive"`
} // @name TestSuiteRequest

// TestRequest represents a test configuration for HTTP requests
// Contains the test definition directly instead of a file reference
type TestRequest struct {
	ID   string     `json:"id" binding:"required" example:"test_1"`
	Test tests.Test `json:"test" binding:"required"`
} // @name TestRequest

// TestExecutionResponse represents the response for test execution
type TestExecutionResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Status    string    `json:"status" example:"running"`
	StartedAt time.Time `json:"started_at" example:"2023-01-01T00:00:00Z"`
	Logs      []string  `json:"logs,omitempty"`
	Error     string    `json:"error,omitempty"`
} // @name TestExecutionResponse

// TestStatusResponse represents the response for test status
type TestStatusResponse struct {
	ID          string     `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Status      string     `json:"status" example:"completed"`
	StartedAt   time.Time  `json:"started_at" example:"2023-01-01T00:00:00Z"`
	CompletedAt *time.Time `json:"completed_at,omitempty" example:"2023-01-01T00:05:00Z"`
	Logs        []string   `json:"logs"`
	Error       string     `json:"error,omitempty"`
} // @name TestStatusResponse

// SystemInfoResponse represents the system information
type SystemInfoResponse struct {
	Version   string `json:"version" example:"1.0.0"`
	GoVersion string `json:"go_version" example:"go1.20.0"`
	OS        string `json:"os" example:"linux"`
	Arch      string `json:"arch" example:"amd64"`
} // @name SystemInfoResponse

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request"`
	Message string `json:"message,omitempty" example:"Detailed error message"`
} // @name ErrorResponse

// TestExecution represents a running test execution
type TestExecution struct {
	ID          string
	Status      string
	StartedAt   time.Time
	CompletedAt *time.Time
	Logs        []string
	Error       string
	mutex       sync.RWMutex
}

// Thread-safe methods for TestExecution
func (te *TestExecution) AddLog(log string) {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	te.Logs = append(te.Logs, log)
}

func (te *TestExecution) SetStatus(status string) {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	te.Status = status
	if status == "completed" || status == "failed" {
		now := time.Now()
		te.CompletedAt = &now
	}
}

func (te *TestExecution) SetError(error string) {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	te.Error = error
	te.Status = "failed"
	now := time.Now()
	te.CompletedAt = &now
}

func (te *TestExecution) GetStatus() TestStatusResponse {
	te.mutex.RLock()
	defer te.mutex.RUnlock()
	return TestStatusResponse{
		ID:          te.ID,
		Status:      te.Status,
		StartedAt:   te.StartedAt,
		CompletedAt: te.CompletedAt,
		Logs:        te.Logs,
		Error:       te.Error,
	}
}

// Global storage for test executions (in production, use a proper database)
var testExecutions = make(map[string]*TestExecution)
var testExecutionsMutex sync.RWMutex

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API server is running
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// ExecuteTestSuite godoc
// @Summary Execute a test suite
// @Description Execute a Voiceflow test suite from request data and return execution ID
// @Tags tests
// @Accept json
// @Produce json
// @Param request body TestExecutionRequest true "Test execution request with embedded suite and tests"
// @Success 202 {object} TestExecutionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tests/execute [post]
func ExecuteTestSuite(c *gin.Context) {
	var req TestExecutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	// Generate unique ID for this execution
	executionID := uuid.New().String()

	// Create test execution record
	execution := &TestExecution{
		ID:        executionID,
		Status:    "running",
		StartedAt: time.Now(),
		Logs:      []string{},
	}

	// Store the execution
	testExecutionsMutex.Lock()
	testExecutions[executionID] = execution
	testExecutionsMutex.Unlock()

	// Start test execution in a goroutine
	go func() {
		// Create a custom logger that captures output
		execution.AddLog("Starting test suite execution...")
		execution.AddLog("Suite: " + req.Suite.Name)
		execution.AddLog("Environment: " + req.Suite.EnvironmentName)

		// Convert the HTTP request to the format expected by the test package
		httpSuite := test.HTTPSuiteRequest{
			Name:               req.Suite.Name,
			Description:        req.Suite.Description,
			EnvironmentName:    req.Suite.EnvironmentName,
			NewSessionPerTest:  req.Suite.NewSessionPerTest,
			Tests:              make([]test.HTTPTestRequest, len(req.Suite.Tests)),
			ApiKey:             req.ApiKey,             // Pass the optional ApiKey
			VoiceflowSubdomain: req.VoiceflowSubdomain, // Pass the optional VoiceflowSubdomain
		}

		// Convert the tests
		for i, testReq := range req.Suite.Tests {
			httpSuite.Tests[i] = test.HTTPTestRequest{
				ID:   testReq.ID,
				Test: testReq.Test,
			}
		}

		// Execute the test suite with log capture
		result := test.ExecuteFromHTTPRequest(httpSuite)

		// Add all captured logs
		for _, logLine := range result.Logs {
			execution.AddLog(logLine)
		}

		if !result.Success {
			execution.SetError(result.Error.Error())
			execution.AddLog("Test suite execution failed: " + result.Error.Error())
		} else {
			execution.SetStatus("completed")
			execution.AddLog("Test suite execution completed successfully")
		}
	}()

	// Return immediate response with execution ID
	response := TestExecutionResponse{
		ID:        executionID,
		Status:    "running",
		StartedAt: execution.StartedAt,
		Logs:      []string{"Test execution started"},
	}

	c.JSON(http.StatusAccepted, response)
}

// GetTestStatus godoc
// @Summary Get test execution status
// @Description Get the status and logs of a test execution
// @Tags tests
// @Accept json
// @Produce json
// @Param id path string true "Test execution ID"
// @Success 200 {object} TestStatusResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/tests/status/{id} [get]
func GetTestStatus(c *gin.Context) {
	executionID := c.Param("id")

	testExecutionsMutex.RLock()
	execution, exists := testExecutions[executionID]
	testExecutionsMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "Test execution not found",
			Message: "No test execution found with ID: " + executionID,
		})
		return
	}

	c.JSON(http.StatusOK, execution.GetStatus())
}

// GetSystemInfo godoc
// @Summary Get system information
// @Description Get information about the API server system
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} SystemInfoResponse
// @Router /api/v1/system/info [get]
func GetSystemInfo(c *gin.Context) {
	response := SystemInfoResponse{
		Version:   global.VersionString,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}

	c.JSON(http.StatusOK, response)
}
