package routes

import (
	"github.com/Dexbooru/dexbooru-core/internal/controllers/misc"
	"github.com/gin-gonic/gin"
)

func registerMiscRoutes(rg *gin.RouterGroup) {
	rg.GET("/health-check", misc.HandleHealthCheck)
}