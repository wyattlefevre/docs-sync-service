package main

import (
	"docs-sync-service/config"
	"docs-sync-service/db"
	"docs-sync-service/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Println("Loaded Config")
	cfg.Print()

	// Set up Gin router
	router := gin.Default()

	db := db.InitDB(cfg)
	defer db.Close()

	// Set up routes
	routes.SetupRoutes(router, db)

	// Start server
	if err := router.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
