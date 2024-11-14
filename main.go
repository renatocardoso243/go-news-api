package main

import (
	"github.com/renatocardoso243/go-news-api.git/database"
	"github.com/renatocardoso243/go-news-api.git/router"
)

func main() {
	// Connect to the database
	database.ConnectDB()
	// Initialize Router
	router.InitializeRouter()
}