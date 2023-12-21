package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllProjects(c *gin.Context) {
	var projects []models.Project

	// Get the user ID
	userIDStr := c.Param("user_id")

	// Convert ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get all projects for the given user ID from the database
	db.DB.Where("user_id = ?", userID).Find(&projects)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": projects})
}
