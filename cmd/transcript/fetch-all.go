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
		startTime, _ := cmd.Flags().GetString("start-time")
		endTime, _ := cmd.Flags().GetString("end-time")
		tag, _ := cmd.Flags().GetString("tag")
		rang, _ := cmd.Flags().GetString("range")

		if err := transcript.FetchAll(agentID, startTime, endTime, tag, rang, outputDirectory); err != nil {
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
	fetchAllCmd.Flags().StringP("start-time", "s", "", "Start time in ISO-8601 format to fetch the analytics. Default is current day but a month ago (optional)")
	fetchAllCmd.Flags().StringP("end-time", "e", "", "Start time in ISO-8601 format to fetch the analytics. Default is current day ago (optional)")
	fetchAllCmd.Flags().StringP("tag", "g", "", "Tag to filter the transcripts. Default is empty (optional)")
	fetchAllCmd.Flags().StringP("range", "r", "", "Range to filter the transcripts. Default is empty (optional)")
	fetchAllCmd.Flags().StringP("output-directory", "d", "./output", "Output directory to save the transcripts. Default is ./output (optional)")

}
