package transcript

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func FetchAll(agentID, startTime, endTime, tag, rang, outputDirectory string) error {
	transcriptsInformation, err := voiceflow.FetchTranscriptInformations(agentID, startTime, endTime, tag, rang)
	if err != nil {
		return fmt.Errorf("failed to fetch transcripts: %w", err)
	}

	for _, transcriptInformation := range transcriptsInformation {
		transcript, err := voiceflow.FetchTranscript(agentID, transcriptInformation.ID)
		if err != nil {
			return fmt.Errorf("failed to fetch transcript %s: %w", transcriptInformation.ID, err)
		}

		err = SaveTranscript(transcript, outputDirectory)
		if err != nil {
			return fmt.Errorf("failed to save transcript %s: %w", transcriptInformation.ID, err)
		}
	}

	return nil
}

func SaveTranscript(transcript [][]string, outputDirectory string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("transcript_%s.csv", timestamp)
	fullPath := filepath.Join(outputDirectory, filename)
	buf := new(bytes.Buffer)
	wr := csv.NewWriter(buf)
	err := wr.WriteAll(transcript)
	if err != nil {
		return fmt.Errorf("failed to write transcript to buffer: %w", err)
	}
	// Write transcript to file
	err = os.WriteFile(fullPath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write transcript file: %w", err)
	}

	return nil
}
