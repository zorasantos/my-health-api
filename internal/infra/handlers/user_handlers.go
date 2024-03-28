package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	err = database.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
