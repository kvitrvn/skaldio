package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/kvitrvn/skaldio/internal/pipeline"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app":     "skaldio",
			"version": "0.1.0",
		})
	})

	pipeline.Mount(r)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
