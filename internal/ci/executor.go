package ci

import (
	"context"
	"strings"
)

// Executor describe the executor architecture
type Executor struct {
	ws WorkspaceInterface
}

// NewExecutor return new executor base on the given workspace
func NewExecutor(ws WorkspaceInterface) *Executor {
	return &Executor{
		ws: ws,
	}
}

// RunDefault execute the pipeline
func (e *Executor) RunDefault(ctx context.Context) (string, error) {
	pipeline, err := e.ws.LoadPipeline()
	if err != nil {
		return "", err
	}

	return e.Run(ctx, *pipeline)
}

// Run the pipeline steps
func (e *Executor) Run(ctx context.Context, pipeline Pipeline) (string, error) {
	output := strings.Builder{}
	output.WriteString("Executing pipeline: ")
	output.WriteString(pipeline.Name)
	output.WriteRune('\n')
	for _, step := range pipeline.Steps {
		output.WriteString("Step: ")
		output.WriteString(step.Name)
		output.WriteRune('\n')
		for _, cmd := range step.Commands {
			withArgs := strings.Fields(cmd)
			cmd := withArgs[:1][0]
			args := withArgs[1:]
			out, err := e.ws.ExecuteCmd(ctx, cmd, args)
			output.Write(out)
			output.WriteRune('\n')
			if err != nil {
				if err := e.ws.Shutdown(); err != nil {
					return "", err
				}
				return output.String(), err
			}
		}
	}

	if err := e.ws.Shutdown(); err != nil {
		return "", err
	}

	return output.String(), nil
}

// Workspace return the workspace attach to the executor
func (e *Executor) Workspace() WorkspaceInterface {
	return e.ws
}
