package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/utils"
)

type ExaminationHandler struct {
	ExaminationDB database.ExaminationInterface
}

func NewExaminationHandler(db database.ExaminationInterface) *ExaminationHandler {
	return &ExaminationHandler{ExaminationDB: db}
}

func (h *ExaminationHandler) CreateExamination(w http.ResponseWriter, r *http.Request) {
	var examination dto.CreateExaminationDTO
	user_id := utils.GetUserIDFromJWT(r)

	err := json.NewDecoder(r.Body).Decode(&examination)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	log.Println("examination dto", examination)

	e, err := entity.NewExamination(user_id, examination.Name, examination.DoctorName, examination.HospitalMedicalRequest, examination.ExaminationDate, examination.Notes)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	err = h.ExaminationDB.Create(e)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Examination created successfully"})

}
