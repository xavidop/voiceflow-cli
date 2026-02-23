package test

import (
	"context"
	"fmt"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

// VoiceflowAgentTestRunner handles the execution of agent-to-agent tests using a Voiceflow agent as the tester
type VoiceflowAgentTestRunner struct {
	*BaseRunner
	// Voiceflow agent tester configuration
	testerEnvironmentName string
	testerAPIKey          string
	testerUserID          string
}

// NewVoiceflowAgentTestRunner creates a new Voiceflow agent test runner
func NewVoiceflowAgentTestRunner(environmentName, userID, apiKeyOverride, subdomainOverride string, logCollector *LogCollector) *VoiceflowAgentTestRunner {
	return &VoiceflowAgentTestRunner{
		BaseRunner:   NewBaseRunner(environmentName, userID, apiKeyOverride, subdomainOverride, logCollector),
		testerUserID: "voiceflow-agent-tester-" + userID, // Different user ID for the tester agent
	}
}

// addLog logs to both the log collector and the global logger for immediate visibility
func (vatr *VoiceflowAgentTestRunner) addLog(message string) {
	vatr.AddLog(message)
}

// ExecuteAgentTest runs an agent-to-agent test using a Voiceflow agent as the tester
func (vatr *VoiceflowAgentTestRunner) ExecuteAgentTest(ctx context.Context, agentTest tests.AgentTest, newSessionPerTest bool) error {
	// Validate that we have the required Voiceflow agent tester configuration
	if agentTest.VoiceflowAgentTesterConfig == nil {
		return fmt.Errorf("voiceflowAgentTesterConfig is required for Voiceflow agent testing")
	}

	// Set up the tester environment configuration
	vatr.testerEnvironmentName = agentTest.VoiceflowAgentTesterConfig.EnvironmentName
	vatr.testerAPIKey = agentTest.VoiceflowAgentTesterConfig.APIKey

	// Configure OpenAI settings for goal evaluation
	vatr.SetOpenAIConfig(agentTest.OpenAIConfig)
	vatr.LogOpenAIConfig()

	vatr.addLog(fmt.Sprintf("Starting Voiceflow agent-to-agent test with goal: %s", agentTest.Goal))
	vatr.addLog(fmt.Sprintf("Target environment: %s", vatr.environmentName))
	vatr.addLog(fmt.Sprintf("Tester environment: %s", vatr.testerEnvironmentName))
	vatr.addLog(fmt.Sprintf("Maximum steps: %d", agentTest.MaxSteps))

	// Initialize both agents
	currentStep := 0
	goalAchieved := false

	var targetAgentResponse []interact.InteractionResponse
	var err error

	// Launch the conversation with the tester agent and provide the goal
	vatr.addLog("Launching conversation with tester Voiceflow agent")
	_, err = vatr.interactWithTesterAgent("launch", "")
	if err != nil {
		return fmt.Errorf("failed to launch conversation with tester agent: %w", err)
	}

	// Update tester agent variables if provided
	if len(agentTest.VoiceflowAgentTesterConfig.Variables) > 0 {
		vatr.addLog(fmt.Sprintf("Setting %d variables in tester agent", len(agentTest.VoiceflowAgentTesterConfig.Variables)))
		err = voiceflow.UpdateStateVariables(
			vatr.testerEnvironmentName,
			vatr.testerUserID,
			agentTest.VoiceflowAgentTesterConfig.Variables,
			vatr.testerAPIKey,
			vatr.subdomainOverride,
		)
		if err != nil {
			return fmt.Errorf("failed to update tester agent variables: %w", err)
		}
		vatr.addLog("Successfully updated tester agent variables")
	}

	// Launch the conversation with the target agent if newSessionPerTest is enabled
	if newSessionPerTest {
		vatr.addLog("Launching conversation with target Voiceflow agent (new session per test enabled)")
		targetAgentResponse, err = vatr.interactWithTargetAgent("launch", "")
		if err != nil {
			return fmt.Errorf("failed to launch conversation with target agent: %w", err)
		}

		// Update target agent variables if provided
		if agentTest.VoiceflowAgentTargetConfig != nil && len(agentTest.VoiceflowAgentTargetConfig.Variables) > 0 {
			vatr.addLog(fmt.Sprintf("Setting %d variables in target agent", len(agentTest.VoiceflowAgentTargetConfig.Variables)))
			err = voiceflow.UpdateStateVariables(
				vatr.environmentName,
				vatr.userID,
				agentTest.VoiceflowAgentTargetConfig.Variables,
				vatr.apiKeyOverride,
				vatr.subdomainOverride,
			)
			if err != nil {
				return fmt.Errorf("failed to update target agent variables: %w", err)
			}
			vatr.addLog("Successfully updated target agent variables")
		}
	}

	// Send the target agent's initial response to the tester agent
	vatr.addLog("Sending target agent's initial response to tester agent")
	// Use the initial response from the target agent as the first message to the tester agent
	// This allows the tester agent to start with the context of the target agent's response
	testerResponse, err := vatr.interactWithTesterAgent("text", vatr.ExtractMessage(targetAgentResponse))
	if err != nil {
		return fmt.Errorf("failed to send goal to tester agent: %w", err)
	}

	currentStep++

	// Main conversation loop
	for currentStep < agentTest.MaxSteps && !goalAchieved {
		// Check for cancellation before each step
		if ctx.Err() != nil {
			vatr.addLog("Agent test execution cancelled")
			return ctx.Err()
		}

		vatr.addLog(fmt.Sprintf("Step %d", currentStep))

		// Get the tester's message and send it to the target agent
		testerMessage := vatr.ExtractMessage(testerResponse)
		if testerMessage == "" {
			vatr.addLog("No message from tester agent, ending conversation")
			break
		}

		// Log the tester agent's message to the target agent if the flag is enabled
		if global.ShowTesterMessages {
			vatr.addLog(fmt.Sprintf("Tester agent says: %s", testerMessage))
		}

		// Send tester's message to target agent
		targetAgentResponse, err = vatr.interactWithTargetAgent("text", testerMessage)
		if err != nil {
			return fmt.Errorf("failed to interact with target agent at step %d: %w", currentStep, err)
		}

		targetMessage := vatr.ExtractMessage(targetAgentResponse)
		vatr.addLog(fmt.Sprintf("Target agent says: %s", targetMessage))

		// Check if goal is achieved by analyzing the conversation
		achieved, err := vatr.IsGoalAchieved(agentTest.Goal)
		if err != nil {
			vatr.addLog(fmt.Sprintf("Error checking goal: %v", err))
		} else if achieved {
			vatr.addLog("Goal achieved successfully!")
			goalAchieved = true
			break
		}

		// Send target agent's response back to tester agent
		testerResponse, err = vatr.interactWithTesterAgent("text", targetMessage)
		if err != nil {
			return fmt.Errorf("failed to get response from tester agent at step %d: %w", currentStep, err)
		}

		currentStep++
	}

	if !goalAchieved {
		// Final goal check
		achieved, err := vatr.IsGoalAchieved(agentTest.Goal)
		if err != nil {
			vatr.addLog(fmt.Sprintf("Error in final goal check: %v", err))
		} else if achieved {
			vatr.addLog("Goal achieved successfully!")
			goalAchieved = true
		}

		if !goalAchieved {
			return fmt.Errorf("goal not achieved within %d steps", agentTest.MaxSteps)
		}
	}

	vatr.addLog(fmt.Sprintf("Voiceflow agent test completed successfully in %d steps", currentStep))
	return nil
}

// interactWithTargetAgent sends a message to the target Voiceflow agent being tested
func (vatr *VoiceflowAgentTestRunner) interactWithTargetAgent(messageType, message string) ([]interact.InteractionResponse, error) {
	responses, err := vatr.InteractWithVoiceflow(messageType, message, vatr.environmentName, vatr.userID, vatr.apiKeyOverride)
	if err != nil {
		return nil, err
	}

	// Add to chat history for goal evaluation
	if messageType != "launch" && message != "" {
		vatr.AddToChatHistory("user", message)
	}

	// Add target agent response to chat history
	targetMessage := vatr.ExtractMessage(responses)
	if targetMessage != "" {
		vatr.AddToChatHistory("assistant", targetMessage)
	}

	return responses, nil
}

// interactWithTesterAgent sends a message to the tester Voiceflow agent
func (vatr *VoiceflowAgentTestRunner) interactWithTesterAgent(messageType, message string) ([]interact.InteractionResponse, error) {
	return vatr.InteractWithVoiceflow(messageType, message, vatr.testerEnvironmentName, vatr.testerUserID, vatr.testerAPIKey)
}
