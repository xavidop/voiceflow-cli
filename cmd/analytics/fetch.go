package analytics

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/analytics"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Aliases: []string{"f", "download", "export"},
	Short:   "Fetch all project analytics. They could write into a file",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		agentID, _ := cmd.Flags().GetString("agent-id")
		outputFile, _ := cmd.Flags().GetString("output-file")
		startTime, _ := cmd.Flags().GetString("start-time")
		endTime, _ := cmd.Flags().GetString("end-time")
		limit, _ := cmd.Flags().GetInt("limit")
		analyticsToFetch, _ := cmd.Flags().GetStringArray("analytics")

		if err := analytics.Fetch(agentID, outputFile, startTime, endTime, limit, analyticsToFetch); err != nil {
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
	analyticsCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringP("agent-id", "a", "", "Voiceflow Agent ID (required)")
	if err := fetchCmd.MarkFlagRequired("agent-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	fetchCmd.Flags().StringP("start-time", "s", "", "Start time in ISO-8601 format to fetch the analytics. Default is current day but a month ago (optional)")
	fetchCmd.Flags().StringP("end-time", "e", "", "Start time in ISO-8601 format to fetch the analytics. Default is current day ago (optional)")
	fetchCmd.Flags().IntP("limit", "l", 100, "Limit of analytics to fetch. Default is 100 (optional)")
	fetchCmd.Flags().StringArrayP("analytics", "t", strings.Split("interactions,sessions,top_intents,top_slots,understood_messages,unique_users,token_usage", ","), "Analytics to fetch. Default is interactions,sessions,top_intents,top_slots,understood_messages,unique_users,token_usage (optional)")
	fetchCmd.Flags().StringP("output-file", "d", "analytics.json", "Output directory to save the analytics. Default is analytics.json (optional)")

}
