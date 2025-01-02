package transcript

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/transcript"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Aliases: []string{"f", "fetch-one", "donwload-one"},
	Short:   "Fetch one transcripts from a project",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		agentID, _ := cmd.Flags().GetString("agent-id")
		transcriptID, _ := cmd.Flags().GetString("transcript-id")
		outputDirectory, _ := cmd.Flags().GetString("output-directory")

		if err := transcript.Fetch(agentID, transcriptID, outputDirectory); err != nil {
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
	transcriptCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringP("agent-id", "a", "", "Voiceflow Agent ID (required)")
	if err := fetchCmd.MarkFlagRequired("agent-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	fetchCmd.Flags().StringP("transcript-id", "t", "", "Voiceflow Transcript ID (required)")
	if err := fetchCmd.MarkFlagRequired("transcript-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	fetchCmd.Flags().StringP("output-directory", "d", "./output", "Output directory to save the transcripts. Default is ./output (optional)")

}
