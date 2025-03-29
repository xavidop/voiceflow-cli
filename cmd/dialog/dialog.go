package dialog

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
)

// DialogCmd represents the dialog command
var dialogCmd = &cobra.Command{
	Use:   "dialog",
	Short: "Start a dialog with the Voiceflow project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(dialogCmd)
}
