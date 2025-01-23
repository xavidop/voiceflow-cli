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

type Payload struct {
	Time    int64       `json:"time,omitempty"`
	Type    string      `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type BlockPayload struct {
	BlockID string `json:"blockID"`
}

type FlowPayload struct {
	DiagramID string `json:"diagramID"`
}

type DebugPayload struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message"`
}

type TextPayload struct {
	Slate   Slate  `json:"slate"`
	Message string `json:"message"`
	Delay   int    `json:"delay"`
	AI      bool   `json:"ai"`
}

type Slate struct {
	ID                       string    `json:"id"`
	Content                  []Content `json:"content"`
	MessageDelayMilliseconds int       `json:"messageDelayMilliseconds"`
}

type Content struct {
	Children []Children `json:"children"`
}

type Children struct {
	Text string `json:"text"`
}

type RequestPayload struct {
	Type    string        `json:"type"`
	Payload IntentPayload `json:"payload"`
}

type IntentPayload struct {
	Query      string        `json:"query"`
	Intent     Intent        `json:"intent"`
	Entities   []interface{} `json:"entities"`
	Confidence float64       `json:"confidence"`
}

type Intent struct {
	Name string `json:"name"`
}

type AIResponseParameters struct {
	System       string  `json:"system"`
	Assistant    string  `json:"assistant"`
	Output       string  `json:"output"`
	Model        string  `json:"model"`
	Temperature  float64 `json:"temperature"`
	MaxTokens    int     `json:"maxTokens"`
	QueryTokens  int     `json:"queryTokens"`
	AnswerTokens int     `json:"answerTokens"`
	Tokens       int     `json:"tokens"`
	Multiplier   float64 `json:"multiplier"`
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

func (p Payload) GetRequestPayload() (*RequestPayload, error) {
	// Modified condition to accept both "request" and "intent" types
	if p.Type != "request" {
		return nil, fmt.Errorf("payload type is not 'request' or 'intent', got: %s", p.Type)
	}

	payloadBytes, err := json.Marshal(p.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	var requestPayload RequestPayload
	if err := json.Unmarshal(payloadBytes, &requestPayload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request payload: %w", err)
	}

	return &requestPayload, nil
}

// Add GetIntentPayload method
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
