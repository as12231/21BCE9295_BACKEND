package handlers

import (
	"file-sharing-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListFiles - List all files uploaded by the user
func ListFiles(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	// Fetch files from DB
	files, err := models.GetFilesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve files"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
