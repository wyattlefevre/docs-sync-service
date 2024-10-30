package routes

import (
	"database/sql"
	"docs-sync-service/handlers"
	"docs-sync-service/repository"
	"docs-sync-service/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	fileRepo := repository.NewFileRepository(db)
	fileService := services.NewFileService(fileRepo)
	fileHandler := handlers.NewFileHandler(fileService)

	api := router.Group("/api/files")
	{
		api.POST("/", fileHandler.Create)
		api.GET("/", fileHandler.GetAll)
	}
}
