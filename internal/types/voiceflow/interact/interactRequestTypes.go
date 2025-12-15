package interact

type EventPayload struct {
	Event EventData `json:"event"`
}

type EventData struct {
	Name string `json:"name"`
}

type IntentPayload struct {
	Intent   IntentData `json:"intent"`
	Entities []Entity   `json:"entities,omitempty"`
}

type IntentData struct {
	Name string `json:"name"`
}

type Entity struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Action struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type Config struct {
	TTS       bool     `json:"tts"`
	StripSSML bool     `json:"stripSSML"`
	StopTypes []string `json:"stopTypes"`
}

type State struct {
	Variables map[string]interface{} `json:"variables"`
}

type InteratctionRequest struct {
	Action Action  `json:"action"`
	Config *Config `json:"config,omitempty"`
	State  *State  `json:"state,omitempty"`
}
