package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	// Get task ID from route parameters
	taskIDStr := c.Param("task_id")

	// Convert task ID to integer type
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Get user ID and project ID from route parameters
	userIDStr := c.Param("user_id")
	projectIDStr := c.Param("project_id")

	// Convert user ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert project ID to integer type
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Delete the task from the database
	result := db.DB.Where("id = ? AND user_id = ? AND project_id = ?", taskID, userID, projectID).Delete(&models.Task{})

	// Check if the task was found and deleted
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
