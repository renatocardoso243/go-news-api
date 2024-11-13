package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/go-news-api.git/database"
)

func main() {
	// Connect to the database
	database.ConnectDB()
	// Initialize gin
	r := gin.Default()

	// Set up routes

	r.Run(":8080")
}