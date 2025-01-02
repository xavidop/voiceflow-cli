package tests

type Suite struct {
	Name            string     `yaml:"name" json:"name"`
	Description     string     `yaml:"description" json:"description"`
	EnvironmentName string     `yaml:"environmentName" json:"environmentName"`
	Tests           []TestFile `yaml:"tests" json:"tests"`
}

type TestFile struct {
	ID   string `yaml:"id" json:"id"`
	File string `yaml:"file" json:"file"`
}
