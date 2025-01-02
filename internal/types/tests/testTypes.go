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
	ID    string `yaml:"id" json:"id"`
	Type  string `yaml:"type" json:"type"`
	Value string `yaml:"value" json:"value"`
}
