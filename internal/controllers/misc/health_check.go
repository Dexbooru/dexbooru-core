package misc

import (
	"os"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(ctx *gin.Context) {
	apiEnv := os.Getenv("API_ENV")

	ctx.JSON(200, gin.H{
		"environment": apiEnv,
		"status":      "healthy",
	})
}
