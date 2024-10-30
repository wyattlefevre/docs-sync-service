package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "docs-sync-service/models"
    "docs-sync-service/services"
)

type FileHandler struct {
    Service *services.FileService
}

func NewFileHandler(service *services.FileService) *FileHandler {
    return &FileHandler{Service: service}
}

func (h *FileHandler) Create(c *gin.Context) {
    var file models.File
    if err := c.ShouldBindJSON(&file); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.Service.Create(&file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, file)
}

func (h *FileHandler) GetAll(c *gin.Context) {
    files, err := h.Service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, files)
}

// Implement other handlers (GetByID, Update, Delete)

