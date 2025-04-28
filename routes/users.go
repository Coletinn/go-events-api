package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Try to bind JSON payload to the user struct
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(
			http.StatusBadRequest, // better use 400 for invalid input
			gin.H{"message": "Invalid request", "error": err.Error()},
		)
		return
	}

	// Save the user in the database
	if err := user.Save(); err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to save user", "error": err.Error()},
		)
		return
	}

	// If successful, return OK
	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}
