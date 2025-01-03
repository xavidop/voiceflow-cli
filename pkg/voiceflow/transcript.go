package voiceflow

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/transcript"
)

func FetchTranscriptInformations(agentID, startTime, endTime, tag, rang string) ([]transcript.TranscriptInformation, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	if startTime != "" {
		startTime = fmt.Sprintf("&startDate=%s", startTime)
	}
	if endTime != "" {
		endTime = fmt.Sprintf("&endDate=%s", endTime)
	}
	if tag != "" {
		tag = fmt.Sprintf("&tag=%s", tag)
	}
	if rang != "" {
		rang = fmt.Sprintf("&range=%s", rang)
	}
	url := fmt.Sprintf("https://api%s.voiceflow.com/v2/transcripts/%s?%s%s%s%s", global.VoiceflowSubdomain, agentID, startTime, endTime, tag, rang)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []transcript.TranscriptInformation{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []transcript.TranscriptInformation{}, err
	}
	var transcriptInformations []transcript.TranscriptInformation
	err = json.Unmarshal([]byte(string(body)), &transcriptInformations)
	if err != nil {
		return []transcript.TranscriptInformation{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return transcriptInformations, nil
}

func FetchTranscript(agentID, transcriptID string) ([][]string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://api%s.voiceflow.com/v2/transcripts/%s/%s/export?format=csv", global.VoiceflowSubdomain, agentID, transcriptID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "text/csv")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return [][]string{}, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// Remove quotes from the beginning and end of the string
	contentString := string(body)[1 : len(string(body))-2]
	// Remove escaped quotes
	contentString = strings.ReplaceAll(contentString, `\"`, `"`)
	// Replace escaped newlines with actual newlines
	c := strings.Split(contentString, "\\r\\n")
	contentString = strings.Join(c, "\r\n")
	contentString = strings.ReplaceAll(contentString, `\n`, ` `)

	// Parse CSV
	content, err := csv.NewReader(strings.NewReader(contentString)).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return content, nil
}
