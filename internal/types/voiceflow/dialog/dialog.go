package dialog

import "github.com/xavidop/voiceflow-cli/internal/types/tests"

type RecordedConversation struct {
	Name         string                `json:"name"`
	Interactions []RecordedInteraction `json:"interactions"`
}

type RecordedInteraction struct {
	ID              string       `json:"id"`
	User            *tests.User  `json:"user"`
	AgentValidation *tests.Agent `json:"agent_validation,omitempty"`
	AgentResponse   []Agent      `json:"agent,omitempty"`
}

type Agent struct {
	Type  string `yaml:"type" json:"type"`
	Value string `yaml:"value" json:"value"`
}
