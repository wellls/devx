package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wellls/devx/internal/api"
	"github.com/wellls/devx/internal/config"
)

func main() {
	//Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Set up API routes
	api.SetupRoutes(router)

	// Start the server
	address := fmt.Sprintf(":%d", cfg.ServerAddress)
	router.Run(address)
}
