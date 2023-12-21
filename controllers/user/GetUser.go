package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Get user ID
	userIDStr := c.Param("user_id")

	// convert ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get the user from the database by their ID
	var user models.User
	result := db.DB.First(&user, userID)

	// Check if the user was found
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": user})
}
