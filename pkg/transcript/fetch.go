package transcript

import (
	"fmt"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func Fetch(agentID, transcriptID, outputDirectory string) error {

	transcript, err := voiceflow.FetchTranscriptCSV(agentID, transcriptID)
	if err != nil {
		return fmt.Errorf("failed to fetch transcript %s: %w", transcriptID, err)
	}

	err = SaveTranscript(transcript, outputDirectory)
	if err != nil {
		return fmt.Errorf("failed to save transcript %s: %w", transcriptID, err)
	}

	return nil
}
