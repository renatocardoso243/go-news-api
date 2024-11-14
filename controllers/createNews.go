package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/database"
	"github.com/renatocardoso243/go-news-api.git/models"
)

func CreateNews(c *gin.Context) {
	var news models.News
	if err := c.BindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	news.PublishedAt = time.Now()
	database.DB.Create(&news)
	c.JSON(http.StatusOK, news)

}