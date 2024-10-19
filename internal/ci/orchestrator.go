package ci

// WorkspaceInterface describe the workspace mandatories methods
type WorkspaceInterface interface {
	Branch() string
	Commit() string
	Dir() string
	Env() []string
	LoadPipeline() (*Pipeline, error)
}
