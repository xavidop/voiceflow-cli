package transcript

import (
	"encoding/json"
	"fmt"
	"time"
)

type Turn struct {
	TurnID    string    `json:"turnID"`
	Format    string    `json:"format"`
	Type      string    `json:"type"`
	Payload   Payload   `json:"payload"`
	StartTime time.Time `json:"startTime"`
}

// GetTranscriptResponse is the wrapper for the v1 get-transcript API response.
type GetTranscriptResponse struct {
	Transcript TranscriptDetail `json:"transcript"`
}

// TranscriptDetail contains the transcript data from the v1 API.
type TranscriptDetail struct {
	ID            string `json:"id"`
	SessionID     string `json:"sessionID"`
	ProjectID     string `json:"projectID"`
	EnvironmentID string `json:"environmentID"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Logs          []Log  `json:"logs"`
}

// Log represents a single log entry in the v1 transcript response.
type Log struct {
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
	CreatedAt string          `json:"createdAt"`
	UpdatedAt string          `json:"updatedAt"`
}

// LogData is used to partially parse the Data field of a Log entry.
type LogData struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type Payload struct {
	Time    int64       `json:"time,omitempty"`
	Type    string      `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type TextPayload struct {
	Message string `json:"message"`
}

type IntentPayload struct {
	Query      string  `json:"query"`
	Confidence float64 `json:"confidence"`
}

// GetTextPayload extracts TextPayload from generic Payload
func (p Payload) GetTextPayload() (*TextPayload, error) {
	if p.Type != "text" {
		return nil, fmt.Errorf("payload type is not 'text', got: %s", p.Type)
	}

	// Convert the interface{} payload to json bytes
	payloadBytes, err := json.Marshal(p.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Unmarshal into TextPayload struct
	var textPayload TextPayload
	if err := json.Unmarshal(payloadBytes, &textPayload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal text payload: %w", err)
	}

	return &textPayload, nil
}

// GetIntentPayload extracts IntentPayload from generic Payload
func (p Payload) GetIntentPayload() (*IntentPayload, error) {
	if p.Type != "intent" {
		return nil, fmt.Errorf("payload type is not 'intent', got: %s", p.Type)
	}

	payloadBytes, err := json.Marshal(p.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal intent payload: %w", err)
	}

	var intentPayload IntentPayload
	if err := json.Unmarshal(payloadBytes, &intentPayload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal intent payload: %w", err)
	}

	return &intentPayload, nil
}
