package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
)

// VersionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get voiceflow-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		// Not check in development
		cmdutils.CheckUpdate(true)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
