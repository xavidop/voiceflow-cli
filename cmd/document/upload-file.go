package document

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/pkg/document"
)

// uploadFileCmd represents the uploadFile command
var uploadFileCmd = &cobra.Command{
	Use:     "upload-file",
	Aliases: []string{"uf", "uplaod-files"},
	Short:   "Upload a dile to a knowledge base",

	Run: func(cmd *cobra.Command, args []string) {
		voiceflow.SetVoiceflowAPIKey()
		fileToUpload, _ := cmd.Flags().GetString("file")
		name, _ := cmd.Flags().GetString("name")
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		maxChunkSize, _ := cmd.Flags().GetInt("max-chunk-size")
		markdownConversion, _ := cmd.Flags().GetBool("markdown-conversion")
		llmGeneratedQ, _ := cmd.Flags().GetBool("llm-generated-q")
		llmPrependContext, _ := cmd.Flags().GetBool("llm-prepend-context")
		llmBasedChunking, _ := cmd.Flags().GetBool("llm-based-chunking")
		llmContentSummarization, _ := cmd.Flags().GetBool("llm-content-summarization")
		tags, _ := cmd.Flags().GetStringArray("tags")

		if err := document.UploadFile(fileToUpload, name, overwrite, maxChunkSize, markdownConversion, llmGeneratedQ, llmPrependContext, llmBasedChunking, llmContentSummarization, tags); err != nil {
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
	documentCmd.AddCommand(uploadFileCmd)
	uploadFileCmd.Flags().StringP("file", "f", "", "File to upload to the knowledge base (required)")
	if err := uploadFileCmd.MarkFlagRequired("file"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	uploadFileCmd.Flags().StringP("name", "n", "", "Name of the document that will be uploaded to the knowledge base (required)")
	if err := uploadFileCmd.MarkFlagRequired("name"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

	uploadFileCmd.Flags().BoolP("overwrite", "w", false, "Overwrite the document if it already exists in the knowledge base. Default is false (optional)")
	uploadFileCmd.Flags().IntP("max-chunk-size", "m", 1000, "Determines how granularly each document is broken up. Default is 1000 (optional)")
	uploadFileCmd.Flags().BoolP("markdown-conversion", "k", false, "Enable HTML to markdown conversion. Default is false (optional)")
	uploadFileCmd.Flags().BoolP("llm-generated-q", "q", false, "If an LLM to generate a question based on the document context and specific chunk, and prepend it to the chunk. Default is false (optional)")
	uploadFileCmd.Flags().BoolP("llm-prepend-context", "p", false, "LLM to generate a context summary based on the document and chunk context, and prepend it to each chunk. Default is false (optional)")
	uploadFileCmd.Flags().BoolP("llm-based-chunking", "g", false, "LLM to determine the optimal chunking of the document content based on semantic similarity and retrieval effectiveness. Default is false (optional)")
	uploadFileCmd.Flags().BoolP("llm-content-summarization", "s", false, "LLM to summarize and rewrite the content, removing unnecessary information and focusing on important parts to optimize for retrieval. Default is false (optional)")
	uploadFileCmd.Flags().StringArrayP("tags", "t", []string{}, "An array of tag labels to attach to a KB document that can be used to filter document eligibility in query retrieval. Default is empty (optional)")

}
