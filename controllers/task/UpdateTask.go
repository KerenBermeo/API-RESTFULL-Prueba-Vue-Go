package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

// UpdateTask updates the information of a task based on the provided parameters in the route.
func UpdateTask(c *gin.Context) {
	// Get user, project, and task IDs from the route parameters
	userIDStr := c.Param("user_id")
	projectIDStr := c.Param("project_id")
	taskIDStr := c.Param("task_id")

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

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	fmt.Println("Before retrieving the existing task")

	// Retrieve the existing task from the database
	var task models.Task
	if err := db.DB.Where("user_id = ? AND project_id = ? AND id = ?", userID, projectID, taskID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	fmt.Println("After retrieving the existing task")

	// Bind the JSON request body to a Task structure
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task information
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description

	// Save the updated task to the database
	if err := db.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the task"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": task})
}
