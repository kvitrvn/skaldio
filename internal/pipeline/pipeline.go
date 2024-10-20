package pipeline

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/kvitrvn/skaldio/internal/ci"
)

type processRequest struct {
	URL    string `json:"url"`
	Branch string `json:"branch"`
}

func process(ctx *gin.Context) {
	body := &processRequest{}

	if err := ctx.Bind(body); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	ws, err := ci.NewWorkspace("./tmp", body.URL, body.Branch)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"repository": body.URL,
			"result":     err.Error(),
		})
		return
	}

	executor := ci.NewExecutor(ws)
	output, err := executor.RunDefault(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"repository":          body.URL,
			"branch":              ws.Branch(),
			"commit":              ws.Commit(),
			"workspace_directory": ws.Dir(),
			"result":              err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"repository":          body.URL,
		"branch":              ws.Branch(),
		"commit":              ws.Commit(),
		"workspace_directory": ws.Dir(),
		"result":              fmt.Sprintf("Successfully executed pipeline.\n\n%s", output),
	})
}
