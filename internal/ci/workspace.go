package ci

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/yaml.v3"
)

// Workspace describe the environment of each run of CI pipeline
type Workspace struct {
	branch string
	commit string
	dir    string
	env    []string
}

// NewWorkspace create and returns a new workspace
func NewWorkspace(root, url, branch string) (*Workspace, error) {
	dir, err := os.MkdirTemp(root, "ci_workspace")
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Depth:             1,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return &Workspace{
		branch: branch,
		commit: ref.Hash().String(),
		dir:    dir,
		env:    []string{},
	}, nil
}

// Branch return the branch name working on
func (ws *Workspace) Branch() string {
	return ws.branch
}

// Commit return the commit SHA working on
func (ws *Workspace) Commit() string {
	return ws.commit
}

// Dir return the dirname working in
func (ws *Workspace) Dir() string {
	return ws.dir
}

// Env return environment variables of the current execution
func (ws *Workspace) Env() []string {
	return ws.env
}

// LoadPipeline checks and parse the yaml CI file
func (ws *Workspace) LoadPipeline() (*Pipeline, error) {
	data, err := os.ReadFile(filepath.Join(ws.dir, "skaldio.yaml"))
	if err != nil {
		return nil, err
	}

	var pipeline Pipeline
	if err := yaml.Unmarshal(data, &pipeline); err != nil {
		return nil, err
	}

	return &pipeline, nil
}

// ExecuteCmd execute the given command with args
func (ws *Workspace) ExecuteCmd(ctx context.Context, cmd string, args []string) ([]byte, error) {
	command := exec.CommandContext(ctx, cmd, args...)
	command.Dir = ws.dir
	command.Env = append(command.Environ(), ws.Env()...)

	return command.CombinedOutput()
}
