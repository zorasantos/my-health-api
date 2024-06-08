package database

import (
	"database/sql"
	"log"

	"github.com/zorasantos/my-health/internal/entity"
)

type Appointment struct {
	DB *sql.DB
}

func NewAppointment(db *sql.DB) *Appointment {
	return &Appointment{DB: db}
}

func (a *Appointment) Create(appointment *entity.Appointment) error {
	_, err := a.DB.Exec("INSERT INTO appointments (id, user_id, doctor_name, hospital_medical_request, appointment_date, specialty, notes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", appointment.ID, appointment.UserID, appointment.DoctorName, appointment.HospitalMedicalRequest, appointment.AppointmentDate, appointment.Specialty, appointment.Notes, appointment.CreatedAt, appointment.UpdatedAt)

	if err != nil {
		return err
	}

	log.Println("Appointment created successfully")

	return err
}
