package test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/PaesslerAG/jsonpath"
	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/pkg/openai"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

// Function to simulate running a test
func runTest(environmentName, userID string, test tests.Test, apiKeyOverride, subdomainOverride string, logCollector *LogCollector, suiteOpenAIConfig *tests.OpenAIConfig) error {
	logCollector.AddLog("Running Test ID: " + test.Name)

	// Check if this is an agent test
	if test.Agent != nil {
		return runAgentTest(environmentName, userID, test, apiKeyOverride, subdomainOverride, logCollector, suiteOpenAIConfig)
	}

	// Original interaction-based test logic
	var availableButtons []tests.Button // Track buttons from previous choice responses
	for _, interaction := range test.Interactions {
		logCollector.AddLog("Interaction ID: " + interaction.ID)
		logCollector.AddLog("\tInteraction Request Type: " + interaction.User.Type)
		if interaction.User.Type == "text" {
			logCollector.AddLog("\tInteraction Request Payload: " + fmt.Sprintf("%v", interaction.User.Text))
		}
		if interaction.User.Type == "event" {
			logCollector.AddLog("\tInteraction Request Event: " + fmt.Sprintf("%v", interaction.User.Event))
		}
		if interaction.User.Type == "intent" {
			logCollector.AddLog("\tInteraction Request Intent: " + fmt.Sprintf("%v", interaction.User.Intent))
		}
		if interaction.User.Type == "launch" {
			logCollector.AddLog("\tInteraction Request Launch")
		}
		if interaction.User.Type == "button" {
			logCollector.AddLog("\tInteraction Request Button: " + interaction.User.Value)
		}

		interactionResponses, err := voiceflow.DialogManagerInteract(environmentName, userID, interaction, apiKeyOverride, subdomainOverride, availableButtons)
		if err != nil {
			return err
		}
		validations := interaction.Agent.Validate
		validations = autoGenerateValidationsIDs(validations)

		// Track which validations have passed across all responses
		remainingValidations := validations

		for _, interactionResponse := range interactionResponses {
			logCollector.AddLog("\tInteraction Response Type: " + interactionResponse.Type)

			// Check if this is a choice response and extract buttons
			if interactionResponse.Type == "choice" {
				if buttonsInterface, ok := interactionResponse.Payload["buttons"]; ok {
					// Convert buttons from interface{} to []tests.Button
					availableButtons = convertToButtons(buttonsInterface)
					logCollector.AddLog(fmt.Sprintf("\tExtracted %d buttons from choice response", len(availableButtons)))
				}
			}

			// Only validate the remaining validations that haven't passed yet
			remainingValidations, err = validateResponse(interactionResponse, remainingValidations, environmentName, userID, apiKeyOverride, subdomainOverride, logCollector)
			if err != nil {
				return err
			}

		}
		if len(remainingValidations) == 0 {
			logCollector.AddLog("All validations passed for Interaction ID: " + interaction.ID)
		} else {
			// Convert to JSON to automatically omit nil/empty fields
			validationsJSON, _ := json.Marshal(remainingValidations)
			return fmt.Errorf("validation failed for Interaction ID: %s, validation: %s", interaction.ID, string(validationsJSON))
		}
	}
	// No errors, test passed
	return nil
}

// runAgentTest executes an agent-to-agent test
func runAgentTest(environmentName, userID string, test tests.Test, apiKeyOverride, subdomainOverride string, logCollector *LogCollector, suiteOpenAIConfig *tests.OpenAIConfig) error {
	logCollector.AddLog("Executing agent-to-agent test: " + test.Name)

	agentTest := *test.Agent
	// Apply suite-level OpenAI configuration if test doesn't have its own config
	if agentTest.OpenAIConfig == nil && suiteOpenAIConfig != nil {
		agentTest.OpenAIConfig = suiteOpenAIConfig
		logCollector.AddLog("Using suite-level OpenAI configuration")
	}

	// Check if this is a Voiceflow agent testing configuration
	if agentTest.VoiceflowAgentTesterConfig != nil {
		logCollector.AddLog("Using Voiceflow agent as the tester")

		// Create Voiceflow agent test runner
		runner := NewVoiceflowAgentTestRunner(environmentName, userID, apiKeyOverride, subdomainOverride, logCollector)

		// Execute the Voiceflow agent test
		return runner.ExecuteAgentTest(agentTest)
	}

	// Default to OpenAI-based agent testing
	logCollector.AddLog("Using OpenAI as the tester")

	// Create OpenAI agent test runner
	runner := NewAgentTestRunner(environmentName, userID, apiKeyOverride, subdomainOverride, logCollector)

	// Execute the agent test
	return runner.ExecuteAgentTest(agentTest)
}

func autoGenerateValidationsIDs(validations []tests.Validation) []tests.Validation {

	for index, validation := range validations {
		if validation.ID == "" {
			validations[index].ID = uuid.New().String()
		}
	}
	return validations

}

func validateResponse(interactionResponse interact.InteractionResponse, validations []tests.Validation, environmentName, userID, apiKeyOverride, subdomainOverride string, logCollector *LogCollector) ([]tests.Validation, error) {
	messageResponse, ok := getNestedValue(interactionResponse.Payload, "message")
	// Ensure payload is of type Speak before accessing its fields
	// Create a slice to store validations that should be kept
	remainingValidations := make([]tests.Validation, 0)
	if ok {
		message := messageResponse.(string)
		logCollector.AddLog("\tInteraction Response Message: " + message)

		for i := 0; i < len(validations); i++ {
			validation := validations[i]
			passed := false
			if validation.Type == "equals" || validation.Type == "exact_match" {
				if message == validation.Value {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with value: " + validation.Value)
					passed = true
				}
			}
			if validation.Type == "contains" {
				if strings.Contains(message, validation.Value) {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with value: " + validation.Value)
					passed = true
				}
			}
			if validation.Type == "regexp" {
				regexString := validation.Value
				compiledRegexp, err := regexp.Compile(regexString)
				if err != nil {
					errorMsg := "Error compiling regexp: " + err.Error()
					logCollector.AddLog(errorMsg)
					return nil, err
				}
				if compiledRegexp.MatchString(message) {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with value: " + validation.Value)
					passed = true
				}
			}
			if validation.Type == "traceType" {
				if interactionResponse.Type == validation.Value {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with value: " + validation.Value)
					passed = true
				}
			}
			if validation.Type == "similarity" {
				if checkSimilarity(message, validation.Values, *validation.SimilarityConfig, logCollector) {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with values: " + fmt.Sprintf("%v", validation.Values) + " and config " + fmt.Sprintf("%v", *validation.SimilarityConfig))
					passed = true
				}
			}

			if validation.Type == "variable" {
				if checkVariableValue(validation, environmentName, userID, apiKeyOverride, subdomainOverride, logCollector) {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with values: " + validation.Value + " and config " + fmt.Sprintf("%v", *validation.VariableConfig))
					passed = true
				}
			}
			if !passed {
				logCollector.AddLog("\tValidation type: " + validation.Type + " FAILED with value: " + validation.Value)
				remainingValidations = append(remainingValidations, validation)
			}
		}
	} else {
		// If there's no message in the response, keep all validations that require a message
		// Only validations that don't require a message (like traceType, variable) should be processed
		for _, validation := range validations {
			switch validation.Type {
			case "variable":
				// Variable validations don't require a message, so check them
				if checkVariableValue(validation, environmentName, userID, apiKeyOverride, subdomainOverride, logCollector) {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with values: " + validation.Value + " and config " + fmt.Sprintf("%v", *validation.VariableConfig))
					// Don't add to remainingValidations since it passed
				} else {
					logCollector.AddLog("\tValidation type: " + validation.Type + " FAILED with value: " + validation.Value)
					remainingValidations = append(remainingValidations, validation)
				}
			case "traceType":
				// TraceType validations can be checked without a message
				if interactionResponse.Type == validation.Value {
					logCollector.AddLog("\tValidation type: " + validation.Type + " PASSED with value: " + validation.Value)
					// Don't add to remainingValidations since it passed
				} else {
					logCollector.AddLog("\tValidation type: " + validation.Type + " FAILED with value: " + validation.Value)
					remainingValidations = append(remainingValidations, validation)
				}
			default:
				// For message-based validations (equals, exact_match, contains, regexp, similarity),
				// if there's no message, the validation fails
				logCollector.AddLog("\tValidation type: " + validation.Type + " FAILED - no message in response")
				remainingValidations = append(remainingValidations, validation)
			}
		}
	}
	return remainingValidations, nil
}

func checkVariableValue(validation tests.Validation, environmentName, userID, apiKeyOverride, subdomainOverride string, logCollector *LogCollector) bool {

	state, err := voiceflow.FetchStateWithOverrides(environmentName, userID, apiKeyOverride, subdomainOverride)
	if err != nil {
		errorMsg := "Error fetching variable state: " + err.Error()
		logCollector.AddLog(errorMsg)
		return false
	}

	var stateValue interface{}
	if validation.VariableConfig.JsonPath != "" {

		// Make sure the jsonpath starts with $
		jsonPathExpr := validation.VariableConfig.JsonPath
		if !strings.HasPrefix(jsonPathExpr, "$") {
			jsonPathExpr = "$" + jsonPathExpr
		}
		variableValue := state.Variables[validation.VariableConfig.Name]

		// Apply JSONPath expression
		stateValue, err = jsonpath.Get(jsonPathExpr, variableValue)
		if err != nil {
			errorMsg := "Error applying JSONPath: " + err.Error()
			logCollector.AddLog(errorMsg)
			return false
		}

	} else {
		// Use the whole state as the value
		stateValue = state.Variables[validation.VariableConfig.Name]
	}

	// Compare the value with the expected value by converting both to strings
	if fmt.Sprint(stateValue) == fmt.Sprint(validation.Value) {
		return true
	} else {
		errorMsg := "Variable value does not match, expected: " + validation.Value + ", got: " + fmt.Sprint(stateValue)
		logCollector.AddLog(errorMsg)
		return false
	}

}

func checkSimilarity(message string, stringsToEvaluate []string, similarityConfig tests.SimilarityConfig, logCollector *LogCollector) bool {
	switch similarityConfig.Provider {
	case "openai":
		similarity, err := openai.OpenAICheckSimilarity(message, stringsToEvaluate, similarityConfig)
		if err != nil {
			errorMsg := "Error checking similarity: " + err.Error()
			logCollector.AddLog(errorMsg)
			return false
		}
		return similarity >= similarityConfig.SimilarityThreshold

	default:
		errorMsg := "Unsupported provider: " + similarityConfig.Provider
		logCollector.AddLog(errorMsg)
		return false
	}
}

// convertToButtons converts interface{} to []tests.Button
func convertToButtons(buttonsInterface interface{}) []tests.Button {
	var buttons []tests.Button
	buttonsArray, ok := buttonsInterface.([]interface{})
	if !ok {
		return buttons
	}

	for _, btnInterface := range buttonsArray {
		btnMap, ok := btnInterface.(map[string]interface{})
		if !ok {
			continue
		}

		button := tests.Button{}
		if name, ok := btnMap["name"].(string); ok {
			button.Name = name
		}

		if requestInterface, ok := btnMap["request"].(map[string]interface{}); ok {
			if reqType, ok := requestInterface["type"].(string); ok {
				button.Request.Type = reqType
			}
			if payload, ok := requestInterface["payload"].(map[string]interface{}); ok {
				button.Request.Payload = payload
			}
		}

		buttons = append(buttons, button)
	}

	return buttons
}

func getNestedValue(data map[string]interface{}, keys ...string) (interface{}, bool) {
	current := data
	for i, key := range keys {
		value, ok := current[key]
		if !ok {
			return nil, false
		}

		// If it's the last key, return the value
		if i == len(keys)-1 {
			return value, true
		}

		// Otherwise, move deeper
		current, ok = value.(map[string]interface{})
		if !ok {
			return nil, false
		}
	}
	return nil, false
}
