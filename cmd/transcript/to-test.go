package transcript

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/transcript"
)

// toTestCmd represents the toTest command
var toTestCmd = &cobra.Command{
	Use:     "to-test",
	Aliases: []string{"to", "to-tests", "test"},
	Short:   "Transforms a transcript into a test",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		agentID, _ := cmd.Flags().GetString("agent-id")
		transcriptID, _ := cmd.Flags().GetString("transcript-id")
		outputFile, _ := cmd.Flags().GetString("output-file")
		testName, _ := cmd.Flags().GetString("test-name")
		testDescription, _ := cmd.Flags().GetString("test-description")

		if err := transcript.ToTest(agentID, transcriptID, outputFile, testName, testDescription); err != nil {
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
	transcriptCmd.AddCommand(toTestCmd)

	toTestCmd.Flags().StringP("agent-id", "a", "", "Voiceflow Agent ID (required)")
	if err := toTestCmd.MarkFlagRequired("agent-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	toTestCmd.Flags().StringP("transcript-id", "t", "", "Voiceflow Transcript ID (required)")
	if err := toTestCmd.MarkFlagRequired("transcript-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	toTestCmd.Flags().StringP("output-file", "d", "test.yaml", "Output file to save the test. Default is test.yaml (optional)")
	toTestCmd.Flags().StringP("test-name", "n", "Test", "Test name (optional)")
	toTestCmd.Flags().StringP("test-description", "e", "Test", "Test description (optional)")

}
