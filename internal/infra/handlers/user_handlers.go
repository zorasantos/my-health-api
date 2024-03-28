package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	u, err := entity.NewUser(user.Username, user.Password, user.Email)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	err = h.UserDB.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})

}
