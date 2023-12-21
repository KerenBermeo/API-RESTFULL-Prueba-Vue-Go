package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetProject(c *gin.Context) {
	// Get the user ID
	userIDStr := c.Param("user_id")
	// Convert user ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get the project ID
	projectIDStr := c.Param("project_id")
	// Convert project ID to integer type
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Get the project from the database by its ID and user ID
	var project models.Project
	result := db.DB.Where("user_id = ? AND id = ?", userID, projectID).First(&project)

	// Check if the project was found
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": project})
}
