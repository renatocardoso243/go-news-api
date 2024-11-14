package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/database"
	"github.com/renatocardoso243/go-news-api.git/models"
)

// Update extisting news
func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News
	if err := database.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "notice not found!"})
		return
	}
	if err := c.BindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&news)
	c.JSON(http.StatusOK, news)
}