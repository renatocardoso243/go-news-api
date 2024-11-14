package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run server
	router.Run(":8080")
}