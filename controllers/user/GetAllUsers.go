package controllers

import (
	"net/http"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Get all users from the database
	db.DB.Find(&users)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": users})
}
