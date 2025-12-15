package dialog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/dialog"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/internal/utils"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
	"gopkg.in/yaml.v3"
)

func Start(userID, environment, recordFile string, saveTest bool) error {
	var conversationForTests dialog.RecordedConversation
	var conversationForRecording dialog.RecordedConversation

	if recordFile != "" {
		conversationForTests = dialog.RecordedConversation{
			Name:         fmt.Sprintf("Recording_%s", time.Now().Format("20060102_150405")),
			Interactions: []dialog.RecordedInteraction{},
		}
		conversationForRecording = dialog.RecordedConversation{
			Name:         fmt.Sprintf("Recording_%s", time.Now().Format("20060102_150405")),
			Interactions: []dialog.RecordedInteraction{},
		}
	}

	// Set up signal handling for Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Create a done channel to signal completion
	done := make(chan bool, 1)

	// If we need to save a recording or test, handle cleanup on interrupt
	if recordFile != "" || saveTest {
		go func() {
			<-signalChan
			global.Log.Infof("Received interrupt signal. Saving conversation before exit...")

			if recordFile != "" {
				if err := saveConversation(recordFile, conversationForRecording); err != nil {
					global.Log.Errorf("Error saving conversation: %v", err)
				} else {
					global.Log.Infof("Conversation saved to %s", recordFile)
				}
			}

			if saveTest {
				testFileName := fmt.Sprintf("test_%s.yaml", time.Now().Format("20060102_150405"))
				if err := saveAsTest(testFileName, conversationForTests); err != nil {
					global.Log.Errorf("Error saving test: %v", err)
					return
				}
				global.Log.Infof("Test saved to %s", testFileName)
			}

			done <- true
			os.Exit(0)
		}()
	} else {
		// If no recording needed, just handle the exit gracefully
		go func() {
			<-signalChan
			global.Log.Infof("Received interrupt signal. Exiting...")
			done <- true
			os.Exit(0)
		}()
	}

	// First, launch the conversation
	global.Log.Infof("Starting conversation with Voiceflow...")
	launchInteraction := tests.Interaction{
		ID: "launch",
		User: tests.User{
			Type: "launch",
		},
	}
	if userID == "" {
		userID = uuid.New().String()
	}

	responses, err := voiceflow.DialogManagerInteract(environment, userID, launchInteraction, "", "", nil)
	if err != nil {
		return fmt.Errorf("error starting dialog: %v", err)
	}

	// Record the launch interaction if needed
	if saveTest {
		// Create a validate array to store responses for tests
		validations := extractValidations(responses)
		recordedInteraction := dialog.RecordedInteraction{
			ID:   launchInteraction.ID,
			User: &launchInteraction.User,
			AgentValidation: &tests.Agent{
				Validate: validations,
			},
		}
		conversationForTests.Interactions = append(conversationForTests.Interactions, recordedInteraction)
	}
	if recordFile != "" {
		// For conversation recording, store just the messages directly
		agentResponses := extractAgentResponses(responses)
		recordedInteraction := dialog.RecordedInteraction{
			ID:            launchInteraction.ID,
			User:          &launchInteraction.User,
			AgentResponse: agentResponses,
		}
		conversationForRecording.Interactions = append(conversationForRecording.Interactions, recordedInteraction)
	}

	// Display initial responses
	displayResponses(responses)

	// Start interactive loop
	reader := bufio.NewReader(os.Stdin)
	interactionCount := 1

	for {
		global.Log.Infof("You: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading input: %v", err)
		}

		// Trim input and check for exit command
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			global.Log.Infof("Ending conversation...")
			break
		}

		// Create interaction for the user input
		interaction := tests.Interaction{
			ID: fmt.Sprintf("interaction_%d", interactionCount),
			User: tests.User{
				Type: "text",
				Text: input,
			},
		}

		// Send user input to Voiceflow
		responses, err := voiceflow.DialogManagerInteract(environment, userID, interaction, "", "", nil)
		if err != nil {
			return fmt.Errorf("error during dialog: %v", err)
		}

		// Display responses
		displayResponses(responses)

		// Record the interaction if needed
		if saveTest {
			// Create a validate array to store responses for tests
			validations := extractValidations(responses)
			recordedInteraction := dialog.RecordedInteraction{
				ID:   interaction.ID,
				User: &interaction.User,
				AgentValidation: &tests.Agent{
					Validate: validations,
				},
			}
			conversationForTests.Interactions = append(conversationForTests.Interactions, recordedInteraction)
		}
		if recordFile != "" {
			// For conversation recording, store just the messages directly
			agentResponses := extractAgentResponses(responses)
			recordedInteraction := dialog.RecordedInteraction{
				ID:            interaction.ID,
				User:          &interaction.User,
				AgentResponse: agentResponses,
			}
			conversationForRecording.Interactions = append(conversationForRecording.Interactions, recordedInteraction)
		}

		interactionCount++
	}

	// Save recorded conversation if recordFile was specified
	if recordFile != "" {
		err := saveConversation(recordFile, conversationForRecording)
		if err != nil {
			return fmt.Errorf("error saving conversation: %v", err)
		}
		global.Log.Infof("Conversation saved to %s", recordFile)
	}

	// Save as test if saveTest flag is true
	if saveTest {
		testFileName := fmt.Sprintf("test_%s.yaml", time.Now().Format("20060102_150405"))
		err := saveAsTest(testFileName, conversationForTests)
		if err != nil {
			return fmt.Errorf("error saving test: %v", err)
		}
		global.Log.Infof("Test saved to %s", testFileName)
	}

	return nil
}

func displayResponses(responses []interact.InteractionResponse) {
	for _, response := range responses {
		switch response.Type {
		case "speak":
			// Extract message from the payload
			if msg, ok := getNestedValue(response.Payload, "message"); ok {
				global.Log.Infof("Voiceflow: %s", msg)
			}
		case "text":
			if msg, ok := getNestedValue(response.Payload, "message"); ok {
				global.Log.Infof("Voiceflow: %s", msg)
			}
		case "visual":
			if msg, ok := getNestedValue(response.Payload, "image"); ok {
				global.Log.Infof("Voiceflow (Image): %s", msg)
			}
		case "end":
			global.Log.Infof("Voiceflow: Conversation ended.")
		default:
			global.Log.Debugf("Unknown response type: %s with payload: %v", response.Type, response.Payload)
		}
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

func saveConversation(filePath string, conversation dialog.RecordedConversation) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer utils.SafeClose(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(conversation)
}

// Extract validations from responses to be saved in the recording
func extractValidations(responses []interact.InteractionResponse) []tests.Validation {
	validations := make([]tests.Validation, 0)

	for _, response := range responses {
		switch response.Type {
		case "speak", "text":
			if msg, ok := getNestedValue(response.Payload, "message"); ok {
				validation := tests.Validation{
					ID:    uuid.NewString(),
					Type:  "equals",
					Value: msg.(string),
				}
				validations = append(validations, validation)
			}
		case "visual":
			if img, ok := getNestedValue(response.Payload, "image"); ok {
				validation := tests.Validation{
					ID:    uuid.NewString(),
					Type:  "equals",
					Value: img.(string),
				}

				validations = append(validations, validation)
			}
		case "end":
			validation := tests.Validation{
				ID:    uuid.NewString(),
				Type:  "traceType",
				Value: "end",
			}

			validations = append(validations, validation)
		default:
			// For other response types, we store the trace type
			validation := tests.Validation{
				ID:    uuid.NewString(),
				Type:  "traceType",
				Value: response.Type,
			}
			validations = append(validations, validation)
		}
	}

	return validations
}

// Extract agent responses from interaction responses
func extractAgentResponses(responses []interact.InteractionResponse) []dialog.Agent {
	agents := []dialog.Agent{}
	messages := []string{}
	hasEnd := false
	otherTypes := []string{}

	// First pass: collect all messages and other types
	for _, response := range responses {
		switch response.Type {
		case "speak", "text":
			if msg, ok := getNestedValue(response.Payload, "message"); ok {
				messages = append(messages, msg.(string))
			}
		case "visual":
			if img, ok := getNestedValue(response.Payload, "image"); ok {
				agents = append(agents, dialog.Agent{
					Type:  "image",
					Value: img.(string),
				})
			}
		case "end":
			hasEnd = true
		default:
			otherTypes = append(otherTypes, response.Type)
		}
	}

	// Add text messages as agents
	for _, msg := range messages {
		agents = append(agents, dialog.Agent{
			Type:  "text",
			Value: msg,
		})
	}

	// Add end flag if present
	if hasEnd {
		agents = append(agents, dialog.Agent{
			Type:  "end",
			Value: "true",
		})
	}

	// Add other types if any
	for _, otherType := range otherTypes {
		agents = append(agents, dialog.Agent{
			Type:  otherType,
			Value: "", // We don't have specific value for these unknown types
		})
	}

	return agents
}

// Convert RecordedConversation to Test and save as YAML file
func saveAsTest(filePath string, conversation dialog.RecordedConversation) error {
	// Create a Test struct from the conversation
	test := tests.Test{
		Name:         conversation.Name,
		Description:  "Generated from dialog recording",
		Interactions: make([]tests.Interaction, 0, len(conversation.Interactions)),
	}

	// Convert RecordedInteraction to Interaction
	for _, recordedInt := range conversation.Interactions {
		// For saved tests, we need to ensure each interaction has an agent with validate field
		interaction := tests.Interaction{
			ID:   recordedInt.ID,
			User: *recordedInt.User,
		}

		// If it's a test interaction, AgentValidation should be populated
		if recordedInt.AgentValidation != nil && len(recordedInt.AgentValidation.Validate) > 0 {
			interaction.Agent = *recordedInt.AgentValidation

			// Add the interaction to the test
			test.Interactions = append(test.Interactions, interaction)
		}
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer utils.SafeClose(file)

	// Encode to YAML
	encoder := yaml.NewEncoder(file)
	defer func() {
		if err := encoder.Close(); err != nil {
			global.Log.Errorf("Error closing YAML encoder: %v", err)
		}
	}()
	return encoder.Encode(test)
}
