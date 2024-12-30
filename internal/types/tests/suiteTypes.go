package tests

type Suite struct {
	Name            string     `yaml:"name"`
	Description     string     `yaml:"description"`
	EnvironmentName string     `yaml:"environmentName"`
	Tests           []TestFile `yaml:"tests"`
}

type TestFile struct {
	ID   string `yaml:"id"`
	File string `yaml:"file"`
}
