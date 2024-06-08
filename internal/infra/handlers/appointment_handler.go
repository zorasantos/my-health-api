package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zorasantos/my-health/internal/dto"
	"github.com/zorasantos/my-health/internal/entity"
	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/utils"
)

type AppointmentHandler struct {
	AppointmentDB database.AppointmentInterface
}

func NewAppointmentHandler(db database.AppointmentInterface) *AppointmentHandler {
	return &AppointmentHandler{AppointmentDB: db}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment dto.CreateAppointmentDTO
	user_id := utils.GetUserIDFromJWT(r)

	err := json.NewDecoder(r.Body).Decode(&appointment)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	a, err := entity.NewAppointment(user_id, appointment.DoctorName, appointment.HospitalMedicalRequest, appointment.AppointmentDate, appointment.Specialty, appointment.Notes)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	err = h.AppointmentDB.Create(a)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]error{"error": err})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Appointment created successfully"})
}
