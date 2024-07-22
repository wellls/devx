package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wellls/devx/internal/api"
	"github.com/wellls/devx/internal/config"
	"github.com/wellls/devx/pkg/database"
)

func main() {
	//Load configuration
	cfg := config.LoadConfig()

	//Start database
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Set up API routes
	api.SetupRoutes(router, db)

	// Start the server
	address := fmt.Sprintf(":%d", cfg.ServerAddress)
	router.Run(address)
}
