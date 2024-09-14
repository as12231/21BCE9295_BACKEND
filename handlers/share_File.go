package handlers

import (
	"file-sharing-system/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShareFile - Generate a public link for a file
func ShareFile(c *gin.Context) {
	fileID := c.Param("file_id")
	// Logic to generate a shareable URL for the file
	shareableLink := storage.GenerateShareableLink(fileID)
	c.JSON(http.StatusOK, gin.H{"url": shareableLink})
}
