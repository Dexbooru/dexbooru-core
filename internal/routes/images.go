package routes

import (
	"github.com/Dexbooru/dexbooru-core/internal/controllers/images"
	"github.com/gin-gonic/gin"
)

func registerImageRoutes(rg *gin.RouterGroup) {	
	rg.POST("/blur", images.HandleBlurImage)
}