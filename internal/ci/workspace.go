package ci

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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
