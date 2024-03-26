package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
)

func CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	u, err := entity.NewUser(user.Username, user.Password, user.Email)

	if err != nil {
		return
	}

	err = database.Create(u)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
