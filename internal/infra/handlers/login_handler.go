package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/utils"
)

func Login(ctx *gin.Context) {
	var user dto.LoginDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	dbUser, _ := database.FindByEmail(user.Email)

	// if err != nil {
	// 	if err.Error() == "error connection db in get user" {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 		return
	// 	}
	// }

	is_match := utils.ComparePasswords(dbUser.Password, user.Password)

	if is_match != nil || dbUser.Email != user.Email {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(dbUser.ID, dbUser.Email, dbUser.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token " + err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
