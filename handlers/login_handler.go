package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zorasantos/my-health/db"
	"github.com/zorasantos/my-health/models"
	"github.com/zorasantos/my-health/utils"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	dbUser, err := db.GetUser(user.Email)

	if err != nil {
		if err.Error() == "error connection db in get user" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	}

	is_match := utils.ComparePasswords(dbUser.Password, user.Password)

	if is_match != nil || dbUser.Email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(dbUser.ID, dbUser.Email, dbUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token " + err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
