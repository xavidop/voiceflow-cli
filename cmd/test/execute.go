package test

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/test"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "execute [suite-path]",
	Aliases: []string{"execute", "e", "exe", "exec"},
	Short:   "Execute a suite",
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		suite := args[0]
		voiceflow.SetVoiceflowAPIKey()
		if err := test.ExecuteSuite(suite); err != nil {
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
	testCmd.AddCommand(executeCmd)

}
