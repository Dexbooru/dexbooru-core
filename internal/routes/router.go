package routes

import "github.com/gin-gonic/gin"

func CreateApplicationRouter() *gin.Engine {
	router := gin.New()

	// define base router groups for the REST API
	v1ApiGroup := router.Group("/api/v1/")	
	imageRouter := v1ApiGroup.Group("/image")
	miscRouter := v1ApiGroup.Group("/misc")

	// register all the subroutes under these router groups
	registerImageRoutes(imageRouter)
	registerMiscRoutes(miscRouter)

	return router
}
