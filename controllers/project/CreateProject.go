package controllers

import (
	"net/http"
	"strconv"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var pj models.Project

	// ShouldBindJSON updates 'pj' variable
	if err := c.ShouldBindJSON(&pj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID
	userIDStr := c.Param("user_id")

	// Convert ID to integer type
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Set the user ID for the project
	pj.UserID = uint(userID)

	// Create the project in the database
	db.DB.Create(&pj)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": pj})
}
