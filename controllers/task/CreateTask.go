package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	// ShouldBindJSON updates 'task' variable
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID and project ID from the route parameters
	userIDStr := c.Param("user_id")
	projectIDStr := c.Param("project_id")

	// Convert IDs to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Set the user ID and project ID for the task
	task.UserID = uint(userID)
	task.ProjectID = uint(projectID)

	// Create the task in the database
	db.DB.Create(&task)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": task})
}
