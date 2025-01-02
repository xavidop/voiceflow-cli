package interact

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
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
	Action Action `json:"action"`
	Config Config `json:"config,omitempty"`
	State  State  `json:"state,omitempty"`
}
