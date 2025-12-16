package tests

type Suite struct {
	Name              string        `yaml:"name" json:"name"`
	Description       string        `yaml:"description" json:"description"`
	EnvironmentName   string        `yaml:"environmentName" json:"environmentName"`
	NewSessionPerTest bool          `yaml:"newSessionPerTest,omitempty" json:"newSessionPerTest,omitempty"`
	Tests             []TestFile    `yaml:"tests" json:"tests"`
	OpenAIConfig      *OpenAIConfig `yaml:"openAIConfig,omitempty" json:"openAIConfig,omitempty"`
}

type TestFile struct {
	ID   string `yaml:"id" json:"id"`
	File string `yaml:"file" json:"file"`
}
