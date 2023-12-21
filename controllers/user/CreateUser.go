package controllers

import (
	"net/http"
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	// ShouldBindJSON updates 'user' variable
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user in the database
	db.DB.Create(&user)

	// Encode the JSON response
	c.JSON(http.StatusOK, gin.H{"data": user})
}
