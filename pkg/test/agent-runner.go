package test

import (
	"fmt"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
)

// AgentTestRunner handles the execution of agent-to-agent tests
type AgentTestRunner struct {
	*BaseRunner
	userInformation map[string]string
}

// ChatMessage represents a message in the conversation history
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// addLog logs to both the log collector and the global logger for immediate visibility
func (atr *AgentTestRunner) addLog(message string) {
	atr.AddLog(message)
}

// NewAgentTestRunner creates a new agent test runner
func NewAgentTestRunner(environmentName, userID, apiKeyOverride, subdomainOverride string, logCollector *LogCollector) *AgentTestRunner {
	return &AgentTestRunner{
		BaseRunner:      NewBaseRunner(environmentName, userID, apiKeyOverride, subdomainOverride, logCollector),
		userInformation: make(map[string]string),
	}
}

// ExecuteAgentTest runs an agent-to-agent test
func (atr *AgentTestRunner) ExecuteAgentTest(agentTest tests.AgentTest) error {
	// Set up user information for easy lookup
	for _, info := range agentTest.UserInformation {
		atr.userInformation[info.Name] = info.Value
	}

	// Configure OpenAI settings for this test
	atr.SetOpenAIConfig(agentTest.OpenAIConfig)
	atr.LogOpenAIConfig()

	atr.addLog(fmt.Sprintf("Starting agent test with goal: %s", agentTest.Goal))
	atr.addLog(fmt.Sprintf("Agent persona: %s", agentTest.Persona))
	atr.addLog(fmt.Sprintf("Maximum steps: %d", agentTest.MaxSteps))

	// Initialize conversation with a system prompt that defines the agent's persona and goal
	systemPrompt := atr.buildSystemPrompt(agentTest)
	atr.AddToChatHistory("system", systemPrompt)

	// Start the conversation by sending a launch event to Voiceflow
	currentStep := 0
	goalAchieved := false

	// Launch the conversation
	atr.addLog("Launching conversation with Voiceflow agent")
	voiceflowResponse, err := atr.interactWithVoiceflow("launch", "")
	if err != nil {
		return fmt.Errorf("failed to launch conversation: %w", err)
	}

	// Process the initial response
	agentResponse, err := atr.getNextAction(voiceflowResponse, agentTest.Goal, currentStep+1, agentTest.MaxSteps)
	if err != nil {
		return fmt.Errorf("failed to get initial action: %w", err)
	}
	currentStep++

	for currentStep < agentTest.MaxSteps && !goalAchieved {
		atr.addLog(fmt.Sprintf("Step %d", currentStep))

		// Check if goal is achieved
		achieved, err := atr.IsGoalAchieved(agentTest.Goal)
		if err != nil {
			atr.addLog(fmt.Sprintf("Error checking goal: %v", err))
		} else if achieved {
			atr.addLog("Goal achieved successfully!")
			goalAchieved = true
			break
		}

		// Check if Voiceflow is requesting user information
		userInfoResponse := atr.checkForUserInfoRequest(voiceflowResponse, agentResponse)
		if userInfoResponse != "" {
			atr.addLog(fmt.Sprintf("Providing user information: %s", userInfoResponse))
			agentResponse = userInfoResponse
		}

		voiceflowResponse, err = atr.interactWithVoiceflow("text", agentResponse)
		if err != nil {
			return fmt.Errorf("failed to interact with Voiceflow at step %d: %w", currentStep, err)
		}

		// Get the next action from the AI agent
		agentResponse, err = atr.getNextAction(voiceflowResponse, agentTest.Goal, currentStep+1, agentTest.MaxSteps)
		if err != nil {
			return fmt.Errorf("failed to get next action at step %d: %w", currentStep, err)
		}
		currentStep++
	}

	if !goalAchieved {
		// Final goal check
		achieved, err := atr.IsGoalAchieved(agentTest.Goal)
		if err != nil {
			atr.addLog(fmt.Sprintf("Error in final goal check: %v", err))
		} else if achieved {
			atr.addLog("Goal achieved successfully!")
			goalAchieved = true
		}

		if !goalAchieved {
			return fmt.Errorf("goal not achieved within %d steps", agentTest.MaxSteps)
		}
	}

	atr.addLog(fmt.Sprintf("Agent test completed successfully in %d steps", currentStep))
	return nil
}

// buildSystemPrompt creates the system prompt for the AI agent
func (atr *AgentTestRunner) buildSystemPrompt(agentTest tests.AgentTest) string {
	userInfoStr := ""
	if len(agentTest.UserInformation) > 0 {
		userInfoList := make([]string, len(agentTest.UserInformation))
		for i, info := range agentTest.UserInformation {
			userInfoList[i] = fmt.Sprintf("- %s: %s", info.Name, info.Value)
		}
		userInfoStr = fmt.Sprintf("\n\nUser Information Available:\n%s", strings.Join(userInfoList, "\n"))
	}

	return fmt.Sprintf(`You are an AI agent testing a conversational AI system. Your goal is to: %s

Your persona: %s

Guidelines:
1. Respond naturally as the persona described above
2. Work towards achieving the stated goal through conversation
3. If the system asks for personal information, use the provided user information when available
4. Keep responses concise and conversational
5. Stay in character throughout the conversation
6. If you encounter requests for information not provided, respond as the persona would (e.g., "I don't have that information")%s

Remember: Your goal is to %s. Work towards this goal while maintaining your persona.`,
		agentTest.Goal,
		agentTest.Persona,
		userInfoStr,
		agentTest.Goal)
}

// interactWithVoiceflow sends a message to the Voiceflow Dialog Manager
func (atr *AgentTestRunner) interactWithVoiceflow(messageType, message string) ([]interact.InteractionResponse, error) {
	return atr.InteractWithVoiceflow(messageType, message, atr.environmentName, atr.userID, atr.apiKeyOverride)
}

// getNextAction uses OpenAI to determine the next action based on the conversation history
func (atr *AgentTestRunner) getNextAction(voiceflowResponse []interact.InteractionResponse, goal string, currentStep, maxSteps int) (string, error) {
	// Extract message from Voiceflow response
	voiceflowMessage := atr.ExtractMessage(voiceflowResponse)
	if voiceflowMessage == "" {
		voiceflowMessage = "No message received"
	}

	// Log the extracted Voiceflow message
	atr.addLog(fmt.Sprintf("Voiceflow message: %s", voiceflowMessage))

	// Add the Voiceflow message to conversation history
	atr.AddToChatHistory("assistant", fmt.Sprintf("Voiceflow said: %s", voiceflowMessage))

	prompt := fmt.Sprintf(`Based on the conversation so far, what should your next response be?

Current step: %d/%d
Goal: %s
Last message from system: %s

Provide only your response message, without any explanation or meta-commentary. Stay in character and work towards your goal.`,
		currentStep, maxSteps, goal, voiceflowMessage)

	// Add the prompt as a user message
	messages := append(atr.GetChatHistory(), ChatMessage{
		Role:    "user",
		Content: prompt,
	})

	// Get response from OpenAI
	atr.addLog("Generating next action using OpenAI...")
	response, err := atr.CallOpenAI(messages)
	if err != nil {
		return "", fmt.Errorf("error generating response: %w", err)
	}

	// Add the agent's response to conversation history
	atr.AddToChatHistory("assistant", response)

	return strings.TrimSpace(response), nil
}

// checkForUserInfoRequest uses OpenAI to determine which user information is being requested and provides it
func (atr *AgentTestRunner) checkForUserInfoRequest(voiceflowResponse []interact.InteractionResponse, agentResponse string) string {
	// Extract message from Voiceflow response
	voiceflowMessage := atr.ExtractMessage(voiceflowResponse)

	// Build available user information list for context
	var availableInfo strings.Builder
	if len(atr.userInformation) > 0 {
		availableInfo.WriteString("Available user information:\n")
		for name, value := range atr.userInformation {
			availableInfo.WriteString(fmt.Sprintf("- %s: %s\n", name, value))
		}
	} else {
		availableInfo.WriteString("No predefined user information available.")
	}

	prompt := fmt.Sprintf(`Analyze the following conversation context to determine if any user information is being requested:

Voiceflow message: "%s"
Agent response: "%s"

%s

Task: Determine if the Voiceflow message or conversation context is requesting any personal information (like name, email, phone, address, account number, etc.).

If a specific piece of information is being requested:
1. If it exists in the available user information, respond with the exact field name (e.g., "email", "name", "phone")
2. If it doesn't exist but is clearly being requested, respond with "INVENT:" followed by the type of information (e.g., "INVENT:email", "INVENT:name", "INVENT:phone")

If no user information is being requested, respond with "NONE".

Response format: Either the field name, "INVENT:type", or "NONE"`,
		voiceflowMessage, agentResponse, availableInfo.String())

	messages := []ChatMessage{
		{Role: "system", Content: "You are a helpful assistant that analyzes conversations to detect when personal information is being requested."},
		{Role: "user", Content: prompt},
	}

	atr.addLog("Analyzing user information request...")
	response, err := atr.CallOpenAI(messages)
	if err != nil {
		atr.addLog(fmt.Sprintf("Error analyzing user info request: %v", err))
		return ""
	}

	response = strings.TrimSpace(response)
	atr.addLog(fmt.Sprintf("User info analysis result: %s", response))

	// Handle the response
	if response == "NONE" {
		return ""
	}

	if strings.HasPrefix(response, "INVENT:") {
		// Need to invent the information
		infoType := strings.TrimPrefix(response, "INVENT:")
		inventedValue := atr.inventUserInformation(infoType)
		atr.addLog(fmt.Sprintf("Invented %s: %s", infoType, inventedValue))
		return inventedValue
	}

	// Check if it's an existing field
	if value, exists := atr.userInformation[response]; exists {
		atr.addLog(fmt.Sprintf("Providing existing user information - %s: %s", response, value))
		return value
	}

	// Fallback: try case-insensitive matching
	lowerResponse := strings.ToLower(response)
	for name, value := range atr.userInformation {
		if strings.ToLower(name) == lowerResponse {
			atr.addLog(fmt.Sprintf("Providing user information (case-insensitive match) - %s: %s", name, value))
			return value
		}
	}

	return ""
}

// inventUserInformation creates realistic user information when it doesn't exist
func (atr *AgentTestRunner) inventUserInformation(infoType string) string {
	// Use OpenAI to generate realistic user information
	prompt := fmt.Sprintf(`Generate realistic %s information for a test user. 

Requirements:
- Make it realistic and believable
- Keep it simple and commonly used
- For emails, use common domains like gmail.com, yahoo.com, outlook.com
- For names, use common first and last names
- For phone numbers, use a realistic format
- For addresses, use realistic city/state combinations

Just respond with the %s value only, no explanation.`, infoType, infoType)

	messages := []ChatMessage{
		{Role: "system", Content: "You are a helpful assistant that generates realistic test data for user information."},
		{Role: "user", Content: prompt},
	}

	response, err := atr.CallOpenAI(messages)
	if err != nil {
		// Fallback to simple defaults if OpenAI fails
		switch strings.ToLower(infoType) {
		case "email":
			return "testuser@example.com"
		case "name", "fullname", "full_name":
			return "John Smith"
		case "firstname", "first_name":
			return "John"
		case "lastname", "last_name":
			return "Smith"
		case "phone", "phonenumber", "phone_number":
			return "555-123-4567"
		case "address":
			return "123 Main St, New York, NY 10001"
		case "account", "accountnumber", "account_number":
			return "ACC123456789"
		default:
			return fmt.Sprintf("test_%s_value", infoType)
		}
	}

	return strings.TrimSpace(response)
}
