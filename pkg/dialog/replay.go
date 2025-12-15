package dialog

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/dialog"
	"github.com/xavidop/voiceflow-cli/internal/utils"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func Replay(userID, environment, recordFile string) error {
	// Check if record file exists
	if recordFile == "" {
		return fmt.Errorf("record file is required")
	}

	// Read the recorded conversation from file
	recordedConversation, err := loadConversation(recordFile)
	if err != nil {
		return fmt.Errorf("error reading recorded conversation: %v", err)
	}

	global.Log.Infof("Replaying conversation: %s", recordedConversation.Name)

	// Set up signal handling for Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Create a done channel to signal completion
	done := make(chan bool, 1)

	// Handle graceful exit
	go func() {
		<-signalChan
		global.Log.Infof("Received interrupt signal. Exiting...")
		done <- true
		os.Exit(0)
	}()

	// Generate a random user ID if not provided
	if userID == "" {
		userID = uuid.New().String()
	}

	// Process each interaction in the recorded conversation
	for i, interaction := range recordedConversation.Interactions {
		// Convert the recorded interaction to a test interaction
		testInteraction := tests.Interaction{
			ID:   interaction.ID,
			User: *interaction.User,
		}

		global.Log.Infof("[%d/%d] Replaying: %s", i+1, len(recordedConversation.Interactions), formatUserInput(interaction.User))

		// Sleep a bit to simulate natural conversation pace
		time.Sleep(500 * time.Millisecond)

		// Send interaction to Voiceflow
		responses, err := voiceflow.DialogManagerInteract(environment, userID, testInteraction, "", "", nil)
		if err != nil {
			return fmt.Errorf("error during dialog: %v", err)
		}

		// Display the responses
		displayResponses(responses)

		// Sleep between interactions to make it more readable
		time.Sleep(1 * time.Second)
	}

	global.Log.Infof("Replay completed successfully.")
	return nil
}

// Read and parse a recorded conversation from file
func loadConversation(filePath string) (dialog.RecordedConversation, error) {
	var conversation dialog.RecordedConversation

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return conversation, err
	}
	defer utils.SafeClose(file)

	// Decode the JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conversation)
	if err != nil {
		return conversation, err
	}

	return conversation, nil
}

// Format user input for display
func formatUserInput(user *tests.User) string {
	switch user.Type {
	case "launch":
		return "Launching conversation"
	case "text":
		return fmt.Sprintf("User said: \"%s\"", user.Text)
	default:
		return fmt.Sprintf("User interaction type: %s", user.Type)
	}
}
