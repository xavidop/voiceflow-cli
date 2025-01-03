package kb

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/kb"
)

// exportCmd represents the query command
var queryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"ask", "queries", "q", "question"},
	Short:   "Query a knowledge base",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		question, _ := cmd.Flags().GetString("question")
		model, _ := cmd.Flags().GetString("model")
		outputFile, _ := cmd.Flags().GetString("output-file")
		temperature, _ := cmd.Flags().GetFloat64("temperature")
		chunkLimit, _ := cmd.Flags().GetInt("chunk-limit")
		synthesis, _ := cmd.Flags().GetBool("synthesis")
		systemPrompt, _ := cmd.Flags().GetString("system-prompt")
		includeTags, _ := cmd.Flags().GetStringArray("include-tags")
		includeOperator, _ := cmd.Flags().GetString("include-operator")
		excludeTags, _ := cmd.Flags().GetStringArray("exclude-tags")
		excludeOperator, _ := cmd.Flags().GetString("exclude-operator")
		includeAllTagged, _ := cmd.Flags().GetBool("include-all-tagged")
		includeAllNonTagged, _ := cmd.Flags().GetBool("include-all-non-tagged")

		if err := kb.Query(question, model, temperature, chunkLimit, synthesis, systemPrompt, includeTags, includeOperator, excludeTags, excludeOperator, includeAllTagged, includeAllNonTagged, outputFile); err != nil {
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
	kbCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringP("question", "q", "", "Question to ask to the knowledge base (required)")
	if err := queryCmd.MarkFlagRequired("question"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	queryCmd.Flags().StringP("model", "m", "", "Model to use while asking the knowledge base (required)")
	if err := queryCmd.MarkFlagRequired("model"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	queryCmd.Flags().Float64P("temperature", "r", 0.7, "Temperature to use while asking the knowledge base. Default to 0.7 (optional)")
	queryCmd.Flags().IntP("chunk-limit", "c", 2, "Chunk limit to use while asking the knowledge base. Default to 2 (optional)")
	queryCmd.Flags().BoolP("synthesis", "s", true, "Indicates whether to use language models to generate an answer. Default to true (optional)")
	queryCmd.Flags().StringP("system-prompt", "p", "", "System prompt to use while asking the knowledge base. Default is empty (optional)")
	queryCmd.Flags().StringArrayP("include-tags", "t", []string{}, "Tags to include. Default is empty (optional)")
	queryCmd.Flags().StringP("include-operator", "i", "", "Tags to include. Possible values: and/or. Default is empty (optional)")
	queryCmd.Flags().StringArrayP("exclude-tags", "y", []string{}, "Tags to exclude. Default is empty (optional)")
	queryCmd.Flags().StringP("exclude-operator", "j", "", "Tags to exclude. Possible values: and/or. Default is empty (optional)")
	queryCmd.Flags().StringP("output-file", "d", "query.json", "Output directory to save the information returned by the CLI. Default is query.json (optional)")
	queryCmd.Flags().BoolP("include-all-tagged", "g", false, "Filters KB documents to include those that have any KB tags attached. Default to false (optional)")
	queryCmd.Flags().BoolP("include-all-non-tagged", "n", false, "Filters KB documents to include those that have no KB tags attached. Default to false (optional)")
}
