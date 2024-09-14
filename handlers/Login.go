package handlers

import (
	"file-sharing-system/utils" // Ensure this matches your module name
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler handles user login and generates JWT
func LoginHandler(c *gin.Context) {
	var loginDetails map[string]string
	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Example: Assume username is provided and valid
	username := loginDetails["username"]

	// Generate JWT token
	token, err := utils.GenerateJWT(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
