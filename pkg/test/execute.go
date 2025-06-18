package test

import (
	"bytes"
	"strings"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

// HTTPSuiteRequest represents a test suite from HTTP request
type HTTPSuiteRequest struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	EnvironmentName string            `json:"environment_name"`
	Tests           []HTTPTestRequest `json:"tests"`
	ApiKey          string            `json:"api_key,omitempty"` // Optional token to override global.VoiceflowAPIKey
}

// HTTPTestRequest represents a test from HTTP request
type HTTPTestRequest struct {
	ID   string     `json:"id"`
	Test tests.Test `json:"test"`
}

// ExecuteFromHTTPRequest executes a test suite directly from HTTP request data
func ExecuteFromHTTPRequest(suiteReq HTTPSuiteRequest) *ExecuteSuiteResult {
	// Create a buffer to capture logs
	var logBuffer bytes.Buffer

	// Save the original output
	originalOutput := global.Log.Out

	// Temporarily redirect the logger output to our buffer
	global.Log.SetOutput(&logBuffer)

	// Execute the test suite
	result := &ExecuteSuiteResult{
		Success: true,
		Logs:    []string{},
	}

	err := executeHTTPSuite(suiteReq)
	if err != nil {
		result.Success = false
		result.Error = err
	}

	// Restore the original output
	global.Log.SetOutput(originalOutput)

	// Parse the logs from the buffer
	logContent := logBuffer.String()
	if logContent != "" {
		// Split by newlines and filter empty lines
		logLines := []string{}
		lines := strings.Split(logContent, "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				logLines = append(logLines, strings.TrimSpace(line))
			}
		}
		result.Logs = logLines
	}

	return result
}

// executeHTTPSuite executes a suite from HTTP request data
func executeHTTPSuite(suiteReq HTTPSuiteRequest) error {
	// Define the user ID
	userID := uuid.New().String()

	global.Log.Infof("Suite: %s\nDescription: %s\nEnvironment: %s\nUser ID: %s",
		suiteReq.Name, suiteReq.Description, suiteReq.EnvironmentName, userID)
	global.Log.Infof("Running Tests:")

	// Execute each test directly from the request data
	for _, testReq := range suiteReq.Tests {
		global.Log.Infof("Running Test ID: %s", testReq.ID)
		err := runTest(suiteReq.EnvironmentName, userID, testReq.Test, suiteReq.ApiKey)
		if err != nil {
			global.Log.Errorf("Error running test %s: %v", testReq.ID, err)
			return err
		}
	}

	return nil
}

func ExecuteSuite(suitesPath string) error {

	// Define the user ID
	userID := uuid.New().String()

	// Load all suites from the path
	suites, err := utils.LoadSuitesFromPath(suitesPath)
	if err != nil {
		global.Log.Errorf("Error loading suites: %v", err)
		return err
	}

	// Iterate over each suite and its tests
	for _, suite := range suites {
		global.Log.Infof("Suite: %s\nDescription: %s\nEnvironment: %s\nUser ID: %s", suite.Name, suite.Description, suite.EnvironmentName, userID)
		global.Log.Infof("Running Tests:")

		for _, testFile := range suite.Tests {
			test, err := utils.LoadTestFromPath(testFile.File)
			if err != nil {
				global.Log.Errorf("Error loading test: %v", err)
				return err
			}
			err = runTest(suite.EnvironmentName, userID, test, "") // No token provided, will use global.VoiceflowAPIKey
			if err != nil {
				global.Log.Errorf("Error running test: %v", err)
				return err
			}
		}
	}
	return nil
}

// ExecuteSuiteResult holds the result of test suite execution including logs
type ExecuteSuiteResult struct {
	Success bool
	Error   error
	Logs    []string
}
