package dialog

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/dialog"
)

// replayCmd represents the replay command
var replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay a dialog with the Voiceflow project",
	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		userID, _ := cmd.Flags().GetString("user-id")
		recordFile, _ := cmd.Flags().GetString("record-file")
		environment, _ := cmd.Flags().GetString("environment")

		// Not check in development
		if err := dialog.Replay(userID, environment, recordFile); err != nil {
			global.Log.Errorf("%s", err.Error())
			os.Exit(1)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	dialogCmd.AddCommand(replayCmd)

	replayCmd.Flags().StringP("user-id", "r", "", "User ID for the dialog (optional)")
	replayCmd.Flags().StringP("record-file", "f", "", "Record file to use (required)")
	replayCmd.Flags().StringP("environment", "e", "development", "Environment to use (optional). Default to development")

}
