package tests

type Test struct {
	Name         string        `yaml:"name" json:"name"`
	Description  string        `yaml:"description" json:"description"`
	Interactions []Interaction `yaml:"interactions,omitempty" json:"interactions,omitempty"`
	Agent        *AgentTest    `yaml:"agent,omitempty" json:"agent,omitempty"`
}

// AgentTest defines an agent-to-agent test configuration
type AgentTest struct {
	Goal                       string                      `yaml:"goal" json:"goal"`
	Persona                    string                      `yaml:"persona" json:"persona"`
	MaxSteps                   int                         `yaml:"maxSteps" json:"maxSteps"`
	UserInformation            []UserInfo                  `yaml:"userInformation,omitempty" json:"userInformation,omitempty"`
	OpenAIConfig               *OpenAIConfig               `yaml:"openAIConfig,omitempty" json:"openAIConfig,omitempty"`
	VoiceflowAgentTesterConfig *VoiceflowAgentTesterConfig `yaml:"voiceflowAgentTesterConfig,omitempty" json:"voiceflowAgentTesterConfig,omitempty"`
}

// VoiceflowAgentTesterConfig defines configuration for using a Voiceflow agent as the tester
type VoiceflowAgentTesterConfig struct {
	EnvironmentName string                 `yaml:"environmentName" json:"environmentName"`
	APIKey          string                 `yaml:"apiKey" json:"apiKey"`
	Variables       map[string]interface{} `yaml:"variables,omitempty" json:"variables,omitempty"`
}

// UserInfo represents information that the agent can use when requested
type UserInfo struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}

type Interaction struct {
	ID      string   `yaml:"id" json:"id"`
	User    User     `yaml:"user" json:"user"`
	Agent   Agent    `yaml:"agent" json:"agent"`
	Buttons []Button `yaml:"-" json:"-"` // Stores buttons from previous response, not serialized
}

type User struct {
	Type   string         `yaml:"type" json:"type"`
	Text   string         `yaml:"text,omitempty" json:"text,omitempty"`
	Value  string         `yaml:"value,omitempty" json:"value,omitempty"` // Used for button label
	Event  string         `yaml:"event,omitempty" json:"event,omitempty"`
	Intent *IntentRequest `yaml:"intent,omitempty" json:"intent,omitempty"`
}

type IntentRequest struct {
	Name     string         `yaml:"name" json:"name"`
	Entities []IntentEntity `yaml:"entities,omitempty" json:"entities,omitempty"`
}

type IntentEntity struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}

// Button represents a button from a choice trace, matching the interact response type
type Button struct {
	Name    string        `json:"name"`
	Request ButtonRequest `json:"request"`
}

// ButtonRequest represents the request object attached to a button
type ButtonRequest struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

type Agent struct {
	Validate []Validation `yaml:"validate" json:"validate"`
}

type Validation struct {
	ID               string            `yaml:"id" json:"id,omitempty"`
	Type             string            `yaml:"type" json:"type"`
	Value            string            `yaml:"value,omitempty" json:"value,omitempty"`
	Values           []string          `yaml:"values,omitempty" json:"values,omitempty"`
	SimilarityConfig *SimilarityConfig `yaml:"similarityConfig,omitempty" json:"similarityConfig,omitempty"`
	VariableConfig   *VariableConfig   `yaml:"variableConfig,omitempty" json:"variableConfig,omitempty"`
}

type VariableConfig struct {
	Name     string `yaml:"name" json:"name"`
	JsonPath string `yaml:"jsonPath,omitempty" json:"jsonPath,omitempty"`
}

type SimilarityConfig struct {
	Provider            string  `yaml:"provider" json:"provider"`
	Model               string  `yaml:"model" json:"model"`
	Temperature         float64 `yaml:"temperature" json:"temperature"`
	SimilarityThreshold float64 `yaml:"similarityThreshold" json:"similarityThreshold"`
	TopK                int     `yaml:"top_k" json:"top_k"`
	TopP                float64 `yaml:"top_p" json:"top_p"`
}

// OpenAIConfig defines OpenAI configuration for agent-to-agent tests
type OpenAIConfig struct {
	Model       string   `yaml:"model,omitempty" json:"model,omitempty"`
	Temperature *float64 `yaml:"temperature,omitempty" json:"temperature,omitempty"`
}
