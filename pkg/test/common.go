package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/internal/utils"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

// BaseRunner provides common functionality for both agent runners
type BaseRunner struct {
	environmentName   string
	userID            string
	apiKeyOverride    string
	subdomainOverride string
	logCollector      *LogCollector
	chatHistory       []ChatMessage
	openAIConfig      *tests.OpenAIConfig
}

// NewBaseRunner creates a new base runner with common functionality
func NewBaseRunner(environmentName, userID, apiKeyOverride, subdomainOverride string, logCollector *LogCollector) *BaseRunner {
	return &BaseRunner{
		environmentName:   environmentName,
		userID:            userID,
		apiKeyOverride:    apiKeyOverride,
		subdomainOverride: subdomainOverride,
		logCollector:      logCollector,
		chatHistory:       make([]ChatMessage, 0),
	}
}

// AddLog logs to both the log collector and the global logger for immediate visibility
func (br *BaseRunner) AddLog(message string) {
	br.logCollector.AddLog(message)
}

// SetOpenAIConfig sets the OpenAI configuration
func (br *BaseRunner) SetOpenAIConfig(config *tests.OpenAIConfig) {
	br.openAIConfig = config
}

// AddToChatHistory adds a message to the chat history
func (br *BaseRunner) AddToChatHistory(role, content string) {
	br.chatHistory = append(br.chatHistory, ChatMessage{
		Role:    role,
		Content: content,
	})
}

// GetChatHistory returns the current chat history
func (br *BaseRunner) GetChatHistory() []ChatMessage {
	return br.chatHistory
}

// InteractWithVoiceflow sends a message to a Voiceflow Dialog Manager
func (br *BaseRunner) InteractWithVoiceflow(messageType, message, environmentName, userID, apiKey string) ([]interact.InteractionResponse, error) {

	// Convert to the expected interaction format
	voiceflowInteraction := tests.Interaction{
		ID: "agent-interaction",
		User: tests.User{
			Type: messageType,
			Text: message,
		},
	}

	// Use the existing Voiceflow interaction method
	responses, err := voiceflow.DialogManagerInteract(environmentName, userID, voiceflowInteraction, apiKey, br.subdomainOverride)
	if err != nil {
		return nil, err
	}

	return br.ProcessResponses(responses), nil
}

// ProcessResponses handles multiple responses by concatenating messages
func (br *BaseRunner) ProcessResponses(responses []interact.InteractionResponse) []interact.InteractionResponse {
	if len(responses) == 0 {
		br.AddLog("No response received from Voiceflow")
		return []interact.InteractionResponse{}
	}

	// If there are multiple responses, concatenate their messages
	if len(responses) > 1 {
		var concatenatedMessage strings.Builder
		for i, response := range responses {
			if message, ok := response.Payload["message"].(string); ok && message != "" {
				if i > 0 {
					concatenatedMessage.WriteString(" ")
				}
				concatenatedMessage.WriteString(message)
			}
		}

		// Update the first response with the concatenated message
		if concatenatedMessage.Len() > 0 {
			responses[0].Payload["message"] = concatenatedMessage.String()
		}

		// Return only the first response with the concatenated message
		return responses[:1]
	}

	return responses
}

// ExtractMessage extracts the message text from Voiceflow response
func (br *BaseRunner) ExtractMessage(voiceflowResponse []interact.InteractionResponse) string {
	if len(voiceflowResponse) > 0 && voiceflowResponse[0].Payload != nil {
		if message, ok := voiceflowResponse[0].Payload["message"].(string); ok {
			return message
		}
	}
	return ""
}

// IsGoalAchieved uses OpenAI to evaluate if the goal has been achieved
func (br *BaseRunner) IsGoalAchieved(goal string) (bool, error) {
	// Build conversation summary
	var conversationSummary strings.Builder
	for _, msg := range br.chatHistory {
		if msg.Role != "system" {
			if msg.Role == "user" {
				conversationSummary.WriteString(fmt.Sprintf("User: %s\n", msg.Content))
			} else if msg.Role == "assistant" {
				conversationSummary.WriteString(fmt.Sprintf("Agent: %s\n", msg.Content))
			} else {
				conversationSummary.WriteString(fmt.Sprintf("%s: %s\n", msg.Role, msg.Content))
			}
		}
	}

	prompt := fmt.Sprintf(`Analyze the following conversation and determine if the goal has been achieved.

Goal: %s

Conversation:
%s

Has the goal been achieved? Respond with only "YES" or "NO".`, goal, conversationSummary.String())

	messages := []ChatMessage{
		{Role: "system", Content: "You are a helpful assistant that analyzes conversations and determines if goals have been achieved."},
		{Role: "user", Content: prompt},
	}

	br.AddLog("Evaluating goal achievement...")
	response, err := br.CallOpenAI(messages)
	if err != nil {
		return false, fmt.Errorf("error evaluating goal: %w", err)
	}

	isAchieved := strings.TrimSpace(strings.ToUpper(response)) == "YES"
	br.AddLog(fmt.Sprintf("Goal achievement evaluation: %s (response: %s)",
		map[bool]string{true: "ACHIEVED", false: "NOT ACHIEVED"}[isAchieved], response))

	return isAchieved, nil
}

// CallOpenAI makes a request to the OpenAI API
func (br *BaseRunner) CallOpenAI(messages []ChatMessage) (string, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Set default values
	model := "gpt-4o"
	temperature := 0.7

	// Override with custom config if available
	if br.openAIConfig != nil {
		// Use custom model if provided, otherwise use default
		if br.openAIConfig.Model != "" {
			model = br.openAIConfig.Model
		}
		// Use custom temperature if explicitly provided (including 0), otherwise use default
		if br.openAIConfig.Temperature != nil {
			temperature = *br.openAIConfig.Temperature
		}
	}

	// Create the request payload
	payload := map[string]interface{}{
		"model":       model,
		"temperature": temperature,
		"messages":    messages,
	}

	// Serialize the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to serialize payload: %w", err)
	}

	// Make the HTTP POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+global.OpenAIAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer utils.SafeClose(resp.Body)

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Extract the response from the API result
	choices := response["choices"].([]interface{})
	if len(choices) == 0 {
		return "", fmt.Errorf("no choices returned in the response")
	}

	responseText := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return responseText, nil
}

// LogOpenAIConfig logs the OpenAI configuration being used
func (br *BaseRunner) LogOpenAIConfig() {
	if br.openAIConfig != nil {
		if br.openAIConfig.Model != "" {
			br.AddLog(fmt.Sprintf("Using OpenAI model: %s", br.openAIConfig.Model))
		}
		if br.openAIConfig.Temperature != nil {
			br.AddLog(fmt.Sprintf("Using OpenAI temperature: %.2f", *br.openAIConfig.Temperature))
		}
	}
}
