package tests

type Test struct {
	Name         string        `yaml:"name"`
	Description  string        `yaml:"description"`
	Interactions []Interaction `yaml:"interactions"`
}

type Interaction struct {
	ID    string `yaml:"id"`
	User  User   `yaml:"user"`
	Agent Agent  `yaml:"agent"`
}

type User struct {
	Type string `yaml:"type"`
	Text string `yaml:"text,omitempty"`
}

type Agent struct {
	Validate []Validation `yaml:"validate"`
}

type Validation struct {
	ID    string `yaml:"id"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}
