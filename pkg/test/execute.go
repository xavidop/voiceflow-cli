package test

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

// HTTPSuiteRequest represents a test suite from HTTP request
type HTTPSuiteRequest struct {
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	EnvironmentName    string              `json:"environment_name"`
	NewSessionPerTest  bool                `json:"new_session_per_test,omitempty"` // Optional flag to create a new user session for each test (default: false)
	Tests              []HTTPTestRequest   `json:"tests"`
	ApiKey             string              `json:"api_key,omitempty"`             // Optional token to override global.VoiceflowAPIKey
	VoiceflowSubdomain string              `json:"voiceflow_subdomain,omitempty"` // Optional subdomain to override global.VoiceflowSubdomain
	OpenAIConfig       *tests.OpenAIConfig `json:"openAIConfig,omitempty"`        // Optional OpenAI configuration for agent tests
}

// HTTPTestRequest represents a test from HTTP request
type HTTPTestRequest struct {
	ID   string     `json:"id"`
	Test tests.Test `json:"test"`
}

// LogCollector is used to collect logs during test execution
type LogCollector struct {
	Logs  []string
	OnLog func(message string) // Optional callback invoked on each log message (e.g. for WebSocket streaming)
}

// AddLog adds a log message to the collector
func (lc *LogCollector) AddLog(message string) {
	// Remove tabs from the message for cleaner logging
	// This is useful to avoid issues with tab characters in logs
	message = strings.ReplaceAll(message, "\t", "")
	lc.Logs = append(lc.Logs, message)
	global.Log.Info(message) // Also log to the global logger
	if lc.OnLog != nil {
		lc.OnLog(message)
	}
}

// ExecuteFromHTTPRequest executes a test suite directly from HTTP request data
func ExecuteFromHTTPRequest(ctx context.Context, suiteReq HTTPSuiteRequest) *ExecuteSuiteResult {
	return ExecuteFromHTTPRequestWithCallback(ctx, suiteReq, nil)
}

// ExecuteFromHTTPRequestWithCallback executes a test suite and invokes onLog for each log line in real-time.
// This is used by the WebSocket handler to stream logs to the client as they happen.
func ExecuteFromHTTPRequestWithCallback(ctx context.Context, suiteReq HTTPSuiteRequest, onLog func(string)) *ExecuteSuiteResult {
	// Create a log collector with optional streaming callback
	logCollector := &LogCollector{
		Logs:  []string{},
		OnLog: onLog,
	}

	// Execute the test suite
	result := &ExecuteSuiteResult{
		Success: true,
		Logs:    []string{},
	}

	err := executeHTTPSuite(ctx, suiteReq, logCollector)
	if err != nil {
		result.Success = false
		result.Error = err
	}

	// Copy logs from collector to result
	result.Logs = logCollector.Logs

	return result
}

// executeHTTPSuite executes a suite from HTTP request data
func executeHTTPSuite(ctx context.Context, suiteReq HTTPSuiteRequest, logCollector *LogCollector) error {
	// Define the user ID
	userID := "test-" + uuid.New().String()

	if suiteReq.VoiceflowSubdomain != "" {
		logCollector.AddLog("Using Voiceflow subdomain: " + suiteReq.VoiceflowSubdomain)
	}

	logCollector.AddLog("Suite: " + suiteReq.Name)
	logCollector.AddLog("Description: " + suiteReq.Description)
	logCollector.AddLog("Environment: " + suiteReq.EnvironmentName)
	if suiteReq.NewSessionPerTest {
		logCollector.AddLog("New session per test: enabled")
	} else {
		logCollector.AddLog("User ID: " + userID)
	}
	logCollector.AddLog("Running Tests:")

	// Execute each test directly from the request data
	for _, testReq := range suiteReq.Tests {
		// Check for cancellation before each test
		if ctx.Err() != nil {
			logCollector.AddLog("Test execution cancelled")
			return ctx.Err()
		}

		// Create a new user ID for each test if newSessionPerTest is enabled
		if suiteReq.NewSessionPerTest {
			userID = "test-" + uuid.New().String()
			logCollector.AddLog("User ID for test " + testReq.ID + ": " + userID)
		}
		logCollector.AddLog("Running Test ID: " + testReq.ID)
		err := runTest(ctx, suiteReq.EnvironmentName, userID, testReq.Test, suiteReq.ApiKey, suiteReq.VoiceflowSubdomain, logCollector, suiteReq.OpenAIConfig, suiteReq.NewSessionPerTest)
		if err != nil {
			errorMsg := "Error running test " + testReq.ID + ": " + err.Error()
			logCollector.AddLog(errorMsg)
			return err
		}
	}

	return nil
}

func ExecuteSuite(suitesPath string) error {

	// Define the user ID
	userID := "test-" + uuid.New().String()

	// Load all suites from the path
	suites, err := utils.LoadSuitesFromPath(suitesPath)
	if err != nil {
		global.Log.Errorf("Error loading suites: %v", err)
		return err
	}

	// Iterate over each suite and its tests
	for _, suite := range suites {
		if suite.NewSessionPerTest {
			global.Log.Infof("Suite: %s\nDescription: %s\nEnvironment: %s\nNew session per test: enabled", suite.Name, suite.Description, suite.EnvironmentName)
		} else {
			global.Log.Infof("Suite: %s\nDescription: %s\nEnvironment: %s\nUser ID: %s", suite.Name, suite.Description, suite.EnvironmentName, userID)
		}
		global.Log.Infof("Running Tests:")

		for _, testFile := range suite.Tests {
			// Create a new user ID for each test if newSessionPerTest is enabled
			if suite.NewSessionPerTest {
				userID = "test-" + uuid.New().String()
				global.Log.Infof("User ID for test %s: %s", testFile.ID, userID)
			}
			test, err := utils.LoadTestFromPath(testFile.File)
			if err != nil {
				global.Log.Errorf("Error loading test: %v", err)
				return err
			}
			// Create a dummy log collector for the existing file-based execution
			logCollector := &LogCollector{Logs: []string{}}
			err = runTest(context.Background(), suite.EnvironmentName, userID, test, "", "", logCollector, suite.OpenAIConfig, suite.NewSessionPerTest) // Pass suite-level OpenAI config
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
