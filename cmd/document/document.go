package document

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
)

// documentCmd represents the document root command
var documentCmd = &cobra.Command{
	Use:     "document",
	Aliases: []string{"doc", "d", "docs", "documents"},
	Short:   "Actions on documents",
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
	rootCmd.AddCommand(documentCmd)
}
