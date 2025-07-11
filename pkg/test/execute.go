package test

import (
	"strings"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

// HTTPSuiteRequest represents a test suite from HTTP request
type HTTPSuiteRequest struct {
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	EnvironmentName    string            `json:"environment_name"`
	Tests              []HTTPTestRequest `json:"tests"`
	ApiKey             string            `json:"api_key,omitempty"`             // Optional token to override global.VoiceflowAPIKey
	VoiceflowSubdomain string            `json:"voiceflow_subdomain,omitempty"` // Optional subdomain to override global.VoiceflowSubdomain
}

// HTTPTestRequest represents a test from HTTP request
type HTTPTestRequest struct {
	ID   string     `json:"id"`
	Test tests.Test `json:"test"`
}

// LogCollector is used to collect logs during test execution
type LogCollector struct {
	Logs []string
}

// AddLog adds a log message to the collector
func (lc *LogCollector) AddLog(message string) {
	// Remove tabs from the message for cleaner logging
	// This is useful to avoid issues with tab characters in logs
	message = strings.ReplaceAll(message, "\t", "")
	lc.Logs = append(lc.Logs, message)
	global.Log.Info(message) // Also log to the global logger
}

// ExecuteFromHTTPRequest executes a test suite directly from HTTP request data
func ExecuteFromHTTPRequest(suiteReq HTTPSuiteRequest) *ExecuteSuiteResult {
	// Create a log collector
	logCollector := &LogCollector{
		Logs: []string{},
	}

	// Execute the test suite
	result := &ExecuteSuiteResult{
		Success: true,
		Logs:    []string{},
	}

	err := executeHTTPSuite(suiteReq, logCollector)
	if err != nil {
		result.Success = false
		result.Error = err
	}

	// Copy logs from collector to result
	result.Logs = logCollector.Logs

	return result
}

// executeHTTPSuite executes a suite from HTTP request data
func executeHTTPSuite(suiteReq HTTPSuiteRequest, logCollector *LogCollector) error {
	// Define the user ID
	userID := uuid.New().String()

	if suiteReq.VoiceflowSubdomain != "" {
		logCollector.AddLog("Using Voiceflow subdomain: " + suiteReq.VoiceflowSubdomain)
	}

	logCollector.AddLog("Suite: " + suiteReq.Name)
	logCollector.AddLog("Description: " + suiteReq.Description)
	logCollector.AddLog("Environment: " + suiteReq.EnvironmentName)
	logCollector.AddLog("User ID: " + userID)
	logCollector.AddLog("Running Tests:")

	// Execute each test directly from the request data
	for _, testReq := range suiteReq.Tests {
		logCollector.AddLog("Running Test ID: " + testReq.ID)
		err := runTest(suiteReq.EnvironmentName, userID, testReq.Test, suiteReq.ApiKey, suiteReq.VoiceflowSubdomain, logCollector)
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
			// Create a dummy log collector for the existing file-based execution
			logCollector := &LogCollector{Logs: []string{}}
			err = runTest(suite.EnvironmentName, userID, test, "", "", logCollector) // No token or subdomain provided, will use global values
			if err != nil {
				global.Log.Errorf("Error running test: %v", err)
				return err
			}
			// Log the collected logs to the global logger for file-based execution
			for _, logLine := range logCollector.Logs {
				global.Log.Info(logLine)
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
