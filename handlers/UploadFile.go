package handlers

import (
	"file-sharing-system/models"
	"file-sharing-system/storage" // Corrected import path
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile - Handles file upload
func UploadFile(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
		return
	}

	// Use local storage to upload the file
	filePath, err := storage.UploadToLocal(file.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Save file metadata in PostgreSQL
	err = models.SaveFileMetadata(userID, file.Filename, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "path": filePath})
}
