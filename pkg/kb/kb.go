package kb

import (
	"fmt"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func Query(question, model string, temperature float64, chunkLimit int, synthesis bool, systemPrompt string, includeTags []string, includeOperator string, excludeTags []string, excludeOperator string, includeAllTagged bool, includeAllNonTagged bool, outputFile string) error {

	query, err := voiceflow.QueryKB(question, model, temperature, chunkLimit, synthesis, systemPrompt, includeTags, includeOperator, excludeTags, excludeOperator, includeAllTagged, includeAllNonTagged)
	if err != nil {
		return fmt.Errorf("failed to query KB: %w", err)
	}

	if outputFile != "" {
		err = voiceflow.SaveQuery(query, outputFile)
		if err != nil {
			return fmt.Errorf("failed to save output: %w", err)
		}
	}

	fmt.Printf("%s\n", query)

	return nil

}
