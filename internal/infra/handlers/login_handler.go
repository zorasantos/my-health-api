package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/utils"
)

type LoginHandler struct {
	UserDB database.UserInterface
}

func UserLoginHandler(db database.UserInterface) *LoginHandler {
	return &LoginHandler{UserDB: db}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	dbUser, _ := h.UserDB.FindByEmail(user.Email)

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]error{"error": err})
	// 	return
	// }

	isMatch := utils.ComparePasswords(dbUser.Password, user.Password)

	if isMatch != nil || dbUser.Email != user.Email {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "%s"}`, `Invalid credentials`)
		return
	}

	token, err := utils.GenerateToken(dbUser.ID, dbUser.Email, dbUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Failed to generate token %s"}`, err.Error())
		return
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"token": "%s"}`, token)
	}
}
