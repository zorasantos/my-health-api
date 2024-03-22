package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-health/models"
	"github.com/my-health/utils"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Login", user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if user.Username != "user" || user.Password != "password" || user.Email != "email" {
		fmt.Println("Login", user)
		token, err := utils.GenerateToken(user.ID, user.Email, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
