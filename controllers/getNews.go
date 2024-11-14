package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/database"
	"github.com/renatocardoso243/go-news-api.git/models"
)

// List all news
func GetNews(c *gin.Context) {
	var news []models.News
	database.DB.Find(&news)
	c.JSON(http.StatusOK, news)	
}