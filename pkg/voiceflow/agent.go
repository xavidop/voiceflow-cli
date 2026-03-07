package voiceflow

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

func ExportAgent(agentID, versionID string) (string, error) {
	url := fmt.Sprintf("%s/v2/versions/%s/export", global.GetAPIBaseURL(""), versionID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)
	req.Header.Add("projectid", agentID)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer utils.SafeClose(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SaveAgent(agent string, outputFile string) error {

	// Write agent to file
	err := os.WriteFile(outputFile, []byte(agent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write agent file: %w", err)
	}

	return nil
}
