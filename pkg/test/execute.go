package test

import (
	"bytes"
	"strings"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

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
			err = runTest(suite.EnvironmentName, userID, test)
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

// ExecuteSuiteWithLogs executes a test suite and captures all logs
func ExecuteSuiteWithLogs(suitesPath string) *ExecuteSuiteResult {
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

	err := ExecuteSuite(suitesPath)
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
