package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/database"
	"github.com/renatocardoso243/go-news-api.git/models"
)

// Delete extisting news
func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.News{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "notice not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notice deleted!"})
}