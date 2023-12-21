package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	// Get the project ID from the route parameters
	pjIDStr := c.Param("project_id")

	// Convert project ID to integer type
	pjID, err := strconv.Atoi(pjIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Get all tasks for the given project from the database
	db.DB.Where("project_id = ?", pjID).Find(&tasks)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetAllTasksWithUserVerification(c *gin.Context) {
	// User verification
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

	// Get all tasks for the given user and project from the database
	var tasks []models.Task
	db.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Find(&tasks)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetAllTasksWithUserVerificationAndErrorHandling(c *gin.Context) {
	// User verification
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

	// Get all tasks for the given user and project from the database
	var tasks []models.Task
	result := db.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Find(&tasks)

	// Check if there was an error during the query
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
		return
	}

	// Check if no tasks were found
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
