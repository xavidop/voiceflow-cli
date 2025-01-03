package agent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/agent"
)

// exportCmd represents the fetch command
var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"e", "download", "fetch"},
	Short:   "Export a voiceflow project into a file",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		agentID, _ := cmd.Flags().GetString("agent-id")
		versionID, _ := cmd.Flags().GetString("version-id")
		outputFile, _ := cmd.Flags().GetString("output-file")

		if err := agent.Export(agentID, versionID, outputFile); err != nil {
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
	agentCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("agent-id", "a", "", "Voiceflow Agent ID (required)")
	if err := exportCmd.MarkFlagRequired("agent-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	exportCmd.Flags().StringP("version-id", "s", "development", "Voiceflow Version ID (optional). Default: development")
	exportCmd.Flags().StringP("output-file", "d", "agent.vf", "Output directory to save the VF file. Default is agent.vf (optional)")

}
