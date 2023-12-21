package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

// UpdateProject updates the information of a project based on the provided parameters in the route.
func UpdateProject(c *gin.Context) {
	// Get user and project IDs from the route parameters
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

	// Retrieve the existing project from the database
	var project models.Project
	if err := db.DB.Where("user_id = ? AND id = ?", userID, projectID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Bind the JSON request body to a Project structure
	var updatedProject models.Project
	if err := c.ShouldBindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the project information
	project.Title = updatedProject.Title

	// Save the updated project to the database
	if err := db.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the project"})
		return
	}

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": project})
}
