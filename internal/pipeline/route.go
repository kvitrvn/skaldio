package pipeline

import "github.com/gin-gonic/gin"

// Mount implement pipelines routes
func Mount(r *gin.Engine) {
	pipelineRoute := r.Group("/p")

	pipelineRoute.POST("/", process)
}
