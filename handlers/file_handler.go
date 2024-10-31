package handlers

import (
	"docs-sync-service/models"
	"docs-sync-service/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	fileId, err := h.Service.Create(&file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fileId)
}

func (h *FileHandler) GetAll(c *gin.Context) {
	docsRepo := c.Query("docs_repo")
	codeRepo := c.Query("code_repo")
	codeFilePath := c.Query("code_file_path")

	if docsRepo != "" && codeRepo != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot filter by both docs_repo and code_repo"})
		return
	}

	if docsRepo != "" {
		files, err := h.Service.GetByDocsRepo(docsRepo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, files)
		return
	}

	if codeRepo != "" && codeFilePath != "" {
		files, err := h.Service.GetByCodeFilePath(codeRepo, codeFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, files)
		return
	}

	if codeRepo != "" {
		files, err := h.Service.GetByCodeRepo(codeRepo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, files)
		return
	}

	files, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Returning files")
	fmt.Printf("%v\n", files)
	c.JSON(http.StatusOK, files)
}

func (h *FileHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	file, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *FileHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file.ID = id // Ensure the file ID is set to the path parameter ID
	if err := h.Service.Update(&file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File updated successfully"})
}

func (h *FileHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
