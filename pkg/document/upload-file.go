package document

import (
	"fmt"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func UploadFile(fileToUpload string, overwrite bool, maxChunkSize int, markdownConversion, llmGeneratedQ, llmPrependContext, llmBasedChunking, llmContentSummarization bool, tags []string) error {

	response, err := voiceflow.UploadDocumentFile(fileToUpload, overwrite, maxChunkSize, markdownConversion, llmGeneratedQ, llmPrependContext, llmBasedChunking, llmContentSummarization, tags)
	if err != nil {
		return fmt.Errorf("failed to query KB: %w", err)
	}

	fmt.Printf("%s\n", response)

	return nil
}
