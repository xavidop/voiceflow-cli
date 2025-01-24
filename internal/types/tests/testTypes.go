package tests

type Test struct {
	Name         string        `yaml:"name" json:"name"`
	Description  string        `yaml:"description" json:"description"`
	Interactions []Interaction `yaml:"interactions" json:"interactions"`
}

type Interaction struct {
	ID    string `yaml:"id" json:"id"`
	User  User   `yaml:"user" json:"user"`
	Agent Agent  `yaml:"agent" json:"agent"`
}

type User struct {
	Type string `yaml:"type" json:"type"`
	Text string `yaml:"text,omitempty" json:"text,omitempty"`
}

type Agent struct {
	Validate []Validation `yaml:"validate" json:"validate"`
}

type Validation struct {
	ID               string            `yaml:"id" json:"id,omitempty"`
	Type             string            `yaml:"type" json:"type"`
	Value            string            `yaml:"value" json:"value"`
	Values           []string          `yaml:"values,omitempty" json:"values,omitempty"`
	SimilarityConfig *SimilarityConfig `yaml:"similarityConfig,omitempty" json:"similarityConfig,omitempty"`
}

type SimilarityConfig struct {
	Provider            string  `yaml:"provider" json:"provider"`
	Model               string  `yaml:"model" json:"model"`
	Temperature         float64 `yaml:"temperature" json:"temperature"`
	SimilarityThreshold float64 `yaml:"similarityThreshold" json:"similarityThreshold"`
	TopK                int     `yaml:"top_k" json:"top_k"`
	TopP                float64 `yaml:"top_p" json:"top_p"`
}
