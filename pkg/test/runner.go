package test

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PaesslerAG/jsonpath"
	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/pkg/openai"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

// Function to simulate running a test
func runTest(environmentName, userID string, test tests.Test) error {
	global.Log.Infof("Running Test ID: %s", test.Name)
	// Here, you would implement the actual test execution logic
	for _, interaction := range test.Interactions {
		global.Log.Infof("Interaction ID: %s", interaction.ID)
		global.Log.Infof("\tInteraction Request Type: %s", interaction.User.Type)
		if interaction.User.Type != "launch" {
			global.Log.Infof("\tInteraction Request Payload: %v", interaction.User.Text)
		}
		interactionResponses, err := voiceflow.DialogManagerInteract(environmentName, userID, interaction)
		if err != nil {
			return err
		}
		validations := interaction.Agent.Validate
		validations = autoGenerateValidationsIDs(validations)

		for _, interactionResponse := range interactionResponses {
			global.Log.Infof("\tInteraction Response Type: %s", interactionResponse.Type)

			validations, err = validateResponse(interactionResponse, validations, environmentName, userID)
			if err != nil {
				return err
			}

		}
		if len(validations) == 0 {
			global.Log.Infof("All validations passed for Interaction ID: %s", interaction.ID)
		} else {
			return fmt.Errorf("validation failed for Interaction ID: %s, not all validations were executed: %v", interaction.ID, validations)
		}
	}
	// No errors, test passed
	return nil
}

func autoGenerateValidationsIDs(validations []tests.Validation) []tests.Validation {

	for index, validation := range validations {
		if validation.ID == "" {
			validations[index].ID = uuid.New().String()
		}
	}
	return validations

}

func validateResponse(interactionResponse interact.InteractionResponse, validations []tests.Validation, environmentName, userID string) ([]tests.Validation, error) {
	messageResponse, ok := getNestedValue(interactionResponse.Payload, "message")
	// Ensure payload is of type Speak before accessing its fields
	// Create a slice to store validations that should be kept
	remainingValidations := make([]tests.Validation, 0)
	if ok {
		message := messageResponse.(string)
		global.Log.Infof("\tInteraction Response Message: %s", message)

		for i := 0; i < len(validations); i++ {
			validation := validations[i]
			passed := false
			if validation.Type == "equals" {
				if message == validation.Value {
					global.Log.Infof("\tValidation type: %s PASSED with value: %s", validation.Type, validation.Value)
					passed = true
				}
			}
			if validation.Type == "contains" {
				if strings.Contains(message, validation.Value) {
					global.Log.Infof("\tValidation type: %s PASSED with value: %s", validation.Type, validation.Value)
					passed = true
				}
			}
			if validation.Type == "regexp" {
				regexString := validation.Value
				compiledRegexp, err := regexp.Compile(regexString)
				if err != nil {
					global.Log.Errorf("Error compiling regexp: %s", err.Error())
					return nil, err
				}
				if compiledRegexp.MatchString(message) {
					global.Log.Infof("\tValidation type: %s PASSED with value: %s", validation.Type, validation.Value)
					passed = true
				}
			}
			if validation.Type == "traceType" {
				if interactionResponse.Type == validation.Value {
					global.Log.Infof("\tValidation type: %s PASSED with value: %s", validation.Type, validation.Value)
					passed = true
				}
			}
			if validation.Type == "similarity" {
				if checkSimilarity(message, validation.Values, *validation.SimilarityConfig) {
					global.Log.Infof("\tValidation type: %s PASSED with values: %v and config %v", validation.Type, validation.Values, *validation.SimilarityConfig)
					passed = true
				}
			}

			if validation.Type == "variable" {
				if checkVariableValue(validation, environmentName, userID) {
					global.Log.Infof("\tValidation type: %s PASSED with values: %v and config %v", validation.Type, validation.Value, *validation.VariableConfig)
					passed = true
				}
			}
			if !passed {
				global.Log.Infof("\tValidation type: %s FAILED with value: %s", validation.Type, validation.Value)
				remainingValidations = append(remainingValidations, validation)
			}
		}
	}
	return remainingValidations, nil
}

func checkVariableValue(validation tests.Validation, environmentName, userID string) bool {

	state, err := voiceflow.FetchState(environmentName, userID)
	if err != nil {
		global.Log.Errorf("Error fetching variable state: %s", err.Error())
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
			global.Log.Errorf("Error applying JSONPath: %v", err)
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
		global.Log.Errorf("Variable value does not match, expected: %s, got: %s", validation.Value, fmt.Sprint(stateValue))
		return false
	}

}

func checkSimilarity(message string, stringsToEvaluate []string, similarityConfig tests.SimilarityConfig) bool {
	switch similarityConfig.Provider {
	case "openai":
		similarity, err := openai.OpenAICheckSimilarity(message, stringsToEvaluate, similarityConfig)
		if err != nil {
			global.Log.Errorf("Error checking similarity: %s", err.Error())
			return false
		}
		return similarity >= similarityConfig.SimilarityThreshold

	default:
		global.Log.Errorf("Unsupported provider: %s", similarityConfig.Provider)
		return false
	}
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
