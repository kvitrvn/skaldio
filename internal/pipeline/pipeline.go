package pipeline

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/kvitrvn/skaldio/internal/ci"
)

type processRequest struct {
	URL string `json:"url"`
}

func process(ctx *gin.Context) {
	body := &processRequest{}

	if err := ctx.Bind(body); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	ws, err := ci.NewWorkspace("./tmp", body.URL, "master")
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"repository":          body.URL,
		"branch":              ws.Branch(),
		"commit":              ws.Commit(),
		"workspace_directory": ws.Dir(),
	})
}
