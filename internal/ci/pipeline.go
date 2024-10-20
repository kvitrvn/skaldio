package ci

// Pipeline describe structure of the ci file
type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

// Step describe each step section of the ci file
type Step struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}
