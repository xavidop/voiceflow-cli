package dialog

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/dialog"
)

// StartCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a dialog with the Voiceflow project",
	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		userID, _ := cmd.Flags().GetString("user-id")
		recordFile, _ := cmd.Flags().GetString("record-file")
		environment, _ := cmd.Flags().GetString("environment")
		saveTest, _ := cmd.Flags().GetBool("save-as-test")

		// Not check in development
		if err := dialog.Start(userID, environment, recordFile, saveTest); err != nil {
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
	dialogCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("user-id", "r", "", "User ID for the dialog (optional)")
	startCmd.Flags().StringP("record-file", "f", "", "Record file to use (optional)")
	startCmd.Flags().StringP("environment", "e", "development", "Environment to use (optional). Default to development")
	startCmd.Flags().BoolP("save-as-test", "t", false, "Save conversation as a test (optional)")
}
