package ci

import "context"

// WorkspaceInterface describe the workspace mandatories methods
type WorkspaceInterface interface {
	Branch() string
	Commit() string
	Dir() string
	Env() []string
	LoadPipeline() (*Pipeline, error)
	ExecuteCmd(ctx context.Context, cmd string, args []string) ([]byte, error)
}

// ExecutorInterface describe the executor mandatories methods
type ExecutorInterface interface {
	Run(ctx context.Context, pipeline Pipeline) (string, error)
	Workspace() WorkspaceInterface
}
