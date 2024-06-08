package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/utils"
)

type PrescriptionHandler struct {
	PrescriptionDB database.PrescriptionInterface
}

func NewPrescriptionHandler(db database.PrescriptionInterface) *PrescriptionHandler {
	return &PrescriptionHandler{PrescriptionDB: db}
}

func (h *PrescriptionHandler) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	var prescription dto.CreatePrescriptionDTO
	user_id := utils.GetUserIDFromJWT(r)

	err := json.NewDecoder(r.Body).Decode(&prescription)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	p, err := entity.NewPrescription(user_id, prescription.MedicamentName, prescription.DoctorName, prescription.PrescriptionDate, prescription.InitialDate, prescription.FinalDate, prescription.Notes, prescription.Dosage)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	err = h.PrescriptionDB.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Prescription created successfully"})
}
