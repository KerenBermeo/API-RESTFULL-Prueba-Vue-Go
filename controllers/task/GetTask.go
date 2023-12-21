package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	userIDStr := c.Param("user_id")
	projectIDStr := c.Param("project_id")
	taskIDStr := c.Param("task_id")

	// Convert user ID to integer type
	userID, errUser := strconv.Atoi(userIDStr)
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert project ID to integer type
	projectID, errProject := strconv.Atoi(projectIDStr)
	if errProject != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Convert task ID to integer type
	taskID, errTask := strconv.Atoi(taskIDStr)
	if errTask != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Get task from the database by its ID, user ID, and project ID
	var task models.Task
	result := db.DB.Where("id = ? AND user_id = ? AND project_id = ?", taskID, userID, projectID).First(&task)

	// Check if the task was found
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": task})
}
