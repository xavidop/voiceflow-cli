package voiceflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/transcript"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

func FetchTranscriptInformations(agentID, startTime, endTime, tag, rang string) ([]transcript.TranscriptInformation, error) {
	const pageSize = 99 // API max take is exclusive < 100
	var allTranscripts []transcript.TranscriptInformation
	seen := make(map[string]bool)

	reqBody := transcript.SearchTranscriptsRequest{}
	if startTime != "" {
		reqBody.StartDate = startTime
	}
	if endTime != "" {
		reqBody.EndDate = endTime
	}

	// Use cursor-based pagination with endDate to avoid skip limits.
	// Results are ordered DESC (newest first), so we move the endDate
	// cursor backwards after each full page.
	for {
		url := fmt.Sprintf("%s/v1/transcript/project/%s?take=%d&skip=0&order=DESC", global.GetAnalyticsBaseURL(), agentID, pageSize)

		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request: %v", err)
		}

		req, err := http.NewRequest("POST", url, strings.NewReader(string(bodyBytes)))
		if err != nil {
			return nil, err
		}

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", global.VoiceflowAPIKey)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		body, err := io.ReadAll(res.Body)
		utils.SafeClose(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
		}

		var searchResponse transcript.SearchTranscriptsResponse
		err = json.Unmarshal(body, &searchResponse)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %v", err)
		}

		if len(searchResponse.Transcripts) == 0 {
			break
		}

		for _, t := range searchResponse.Transcripts {
			if !seen[t.ID] {
				seen[t.ID] = true
				allTranscripts = append(allTranscripts, t)
			}
		}

		// If we got fewer results than the page size, we've reached the end
		if len(searchResponse.Transcripts) < pageSize {
			break
		}

		// Use the oldest transcript's createdAt as the new endDate cursor
		oldest := searchResponse.Transcripts[len(searchResponse.Transcripts)-1]
		reqBody.EndDate = oldest.CreatedAt
	}

	return allTranscripts, nil
}

func FetchTranscriptCSV(agentID, transcriptID string) ([][]string, error) {
	turns, err := FetchTranscriptJSON(agentID, transcriptID)
	if err != nil {
		return nil, err
	}

	// Header row
	rows := [][]string{{"type", "payload_type", "message", "timestamp"}}

	for _, turn := range turns {
		message := ""
		switch turn.Payload.Type {
		case "text":
			if turn.Type == "user-text" {
				// User text payload is a plain string
				if s, ok := turn.Payload.Payload.(string); ok {
					message = s
				}
			} else {
				if tp, err := turn.Payload.GetTextPayload(); err == nil {
					message = tp.Message
				}
			}
		case "intent":
			if ip, err := turn.Payload.GetIntentPayload(); err == nil {
				message = ip.Query
			}
		}

		rows = append(rows, []string{
			turn.Type,
			turn.Payload.Type,
			message,
			turn.StartTime.Format(time.RFC3339),
		})
	}

	return rows, nil
}

func FetchTranscriptJSON(agentID, transcriptID string) ([]transcript.Turn, error) {
	url := fmt.Sprintf("%s/v1/transcript/%s", global.GetAnalyticsBaseURL(), transcriptID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer utils.SafeClose(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	var response transcript.GetTranscriptResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return logsToTurns(response.Transcript.Logs)
}

// logsToTurns converts the v1 API log entries to the Turn slice used by downstream code.
// Action logs represent user input; trace logs with type "text" represent agent responses.
func logsToTurns(logs []transcript.Log) ([]transcript.Turn, error) {
	var turns []transcript.Turn

	for _, log := range logs {
		var data transcript.LogData
		if err := json.Unmarshal(log.Data, &data); err != nil {
			continue
		}

		createdAt, _ := time.Parse(time.RFC3339, log.CreatedAt)

		switch log.Type {
		case "action":
			turn := transcript.Turn{
				StartTime: createdAt,
			}
			switch data.Type {
			case "launch":
				turn.Type = "launch"
			case "intent":
				turn.Type = "request"
				turn.Payload = transcript.Payload{
					Type: "intent",
				}
				if data.Payload != nil {
					var payloadData interface{}
					_ = json.Unmarshal(data.Payload, &payloadData)
					turn.Payload.Payload = payloadData
				}
			case "text":
				// User text input — payload is a plain string
				turn.Type = "user-text"
				var userText string
				if data.Payload != nil {
					_ = json.Unmarshal(data.Payload, &userText)
				}
				turn.Payload = transcript.Payload{
					Type:    "text",
					Payload: userText,
				}
			default:
				continue
			}
			turns = append(turns, turn)

		case "trace":
			if data.Type != "text" {
				continue
			}
			var payloadData interface{}
			if data.Payload != nil {
				_ = json.Unmarshal(data.Payload, &payloadData)
			}
			turns = append(turns, transcript.Turn{
				Type: "text",
				Payload: transcript.Payload{
					Type:    "text",
					Payload: payloadData,
				},
				StartTime: createdAt,
			})
		}
	}

	return turns, nil
}
