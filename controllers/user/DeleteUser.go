package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Get user ID from route parameters
	userIDStr := c.Param("user_id")

	// Convert user ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get all projects for the given user from the database
	var projects []models.Project
	db.DB.Where("user_id = ?", userID).Find(&projects)

	// Iterate over projects and delete associated tasks
	for _, project := range projects {
		var tasks []models.Task
		db.DB.Where("user_id = ? AND project_id = ?", userID, project.ID).Find(&tasks)
		db.DB.Delete(&tasks)
	}

	// Delete all projects associated with the user
	db.DB.Delete(&projects)

	// Delete the user from the database
	result := db.DB.Where("id = ?", userID).Delete(&models.User{})

	// Check if the user was found and deleted
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "User, projects, and associated tasks deleted successfully"})
}
