package main

import (
	"example.com/event-booking-restapi/config"
	"example.com/event-booking-restapi/internal/database"
	"example.com/event-booking-restapi/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()
	
	database.InitDB()
	server := gin.Default()

	// Register routes for different features
	router.RegisterRoutes(server)

	server.Run(":8080")
}