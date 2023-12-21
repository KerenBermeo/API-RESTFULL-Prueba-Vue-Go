package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	// Get user ID from route parameters
	userIDStr := c.Param("user_id")

	// Convert user ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get existing user from the database
	var existingUser models.User
	if err := db.DB.First(&existingUser, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind the JSON request body to a User struct
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user information
	existingUser.Name = updatedUser.Name
	existingUser.Password = updatedUser.Password

	// Save the updated user to the database
	if err := db.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": existingUser})
}
