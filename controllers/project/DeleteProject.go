package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func DeleteProject(c *gin.Context) {
	// Get user ID and project ID from route parameters
	userIDStr := c.Param("user_id")
	projectIDStr := c.Param("project_id")

	// Convert user ID and project ID to integer type
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

	// Delete associated tasks for the project
	if err := DeleteAssociatedTasks(userID, projectID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting associated tasks"})
		return
	}

	// Delete the project from the database
	result := db.DB.Where("user_id = ? AND id = ?", userID, projectID).Delete(&models.Project{})

	// Check if the project was found and deleted
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Project and associated tasks deleted successfully"})
}

// Function to delete associated tasks for a project
func DeleteAssociatedTasks(userID, projectID int) error {
	// Get all tasks for the given user and project from the database
	var tasks []models.Task
	if err := db.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Find(&tasks).Error; err != nil {
		return err
	}

	// Delete all tasks associated with the project
	for _, task := range tasks {
		if err := db.DB.Delete(&task).Error; err != nil {
			return err
		}
	}

	return nil
}
