package transcript

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/transcript"
)

// fetchAllCmd represents the fetchAll command
var fetchAllCmd = &cobra.Command{
	Use:     "fetch-all",
	Aliases: []string{"fa", "fetch-everything", "donwload-all"},
	Short:   "Fetch all transcripts from a project",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		agentID, _ := cmd.Flags().GetString("agent-id")
		outputDirectory, _ := cmd.Flags().GetString("output-directory")

		if err := transcript.FetchAll(agentID, outputDirectory); err != nil {
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
	transcriptCmd.AddCommand(fetchAllCmd)

	fetchAllCmd.Flags().StringP("agent-id", "a", "", "Voiceflow Agent ID (required)")
	if err := fetchAllCmd.MarkFlagRequired("agent-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	fetchAllCmd.Flags().StringP("output-directory", "d", "./output", "Output directory to save the transcripts. Default is ./output (optional)")

}
