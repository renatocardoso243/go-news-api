package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/controllers"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/news", controllers.GetNews)
	router.POST("/news", controllers.CreateNews)
	// Add route to get news by id
	router.PUT("/news/:id", controllers.UpdateNews)
	router.DELETE("/news/:id", controllers.DeleteNews)
}