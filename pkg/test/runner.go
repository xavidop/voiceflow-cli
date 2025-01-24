package test

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/pkg/openai"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

// Function to simulate running a test
func runTest(EnvironmentName, userID string, test tests.Test) error {
	global.Log.Infof("Running Test ID: %s", test.Name)
	// Here, you would implement the actual test execution logic
	for _, interaction := range test.Interactions {
		global.Log.Infof("Interaction ID: %s", interaction.ID)
		global.Log.Infof("\tInteraction Request Type: %s", interaction.User.Type)
		if interaction.User.Type != "launch" {
			global.Log.Infof("\tInteraction Request Payload: %v", interaction.User.Text)
		}
		interactionResponses, err := voiceflow.CallInteractionAPI(EnvironmentName, userID, interaction)
		if err != nil {
			return err
		}
		validations := interaction.Agent.Validate
		validations = autoGenerateValidationsIDs(validations)

		for _, interactionResponse := range interactionResponses {
			global.Log.Infof("\tInteraction Response Type: %s", interactionResponse.Type)

			validations, err = validateResponse(interactionResponse, validations)
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

func validateResponse(interactionResponse interact.InteractionResponse, validations []tests.Validation) ([]tests.Validation, error) {
	messageResponse, ok := getNestedValue(interactionResponse.Payload, "message")
	// Ensure payload is of type Speak before accessing its fields
	if ok {
		message := messageResponse.(string)
		global.Log.Infof("\tInteraction Response Message: %s", message)
		for _, validation := range validations {
			if validation.Type == "equals" {
				if message == validation.Value {
					validations = removeById(validations, validation.ID)
					continue
				}
			}
			if validation.Type == "contains" {
				if strings.Contains(message, validation.Value) {
					validations = removeById(validations, validation.ID)
					continue
				}
			}
			if validation.Type == "regexp" {
				regexString := validation.Value
				compiledRegexp, err := regexp.Compile(regexString)
				if err != nil {
					return validations, err
				}
				if compiledRegexp.MatchString(message) {
					validations = removeById(validations, validation.ID)
					continue
				}
			}
			if validation.Type == "traceType" {
				if interactionResponse.Type == validation.Value {
					validations = removeById(validations, validation.ID)
					continue
				}
			}
			if validation.Type == "similarity" {
				if checkSimilarity(message, validation.Values, *validation.SimilarityConfig) {
					validations = removeById(validations, validation.ID)
					continue
				}
			}
		}
	}
	return validations, nil
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

func removeById(slice []tests.Validation, ID string) []tests.Validation {
	for index, validation := range slice {
		if validation.ID == ID {
			return append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
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
