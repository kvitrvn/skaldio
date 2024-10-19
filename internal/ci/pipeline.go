package ci

// Pipeline describe structure of the ci file
type Pipeline struct {
	Name string `yaml:"name"`
	Step []Step `yaml:"steps"`
}

// Step describe each step section of the ci file
type Step struct {
	Name      string   `yaml:"name"`
	Commandes []string `yaml:"commands"`
}
