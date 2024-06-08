package entity

import (
	"time"

	"github.com/zorasantos/my-health/utils"
)

type Appointment struct {
	ID                     utils.ID   `json:"id"`
	UserID                 string     `json:"user_id"`
	DoctorName             string     `json:"doctor_name"`
	HospitalMedicalRequest string     `json:"hospital_medical_request"`
	AppointmentDate        string     `json:"appointment_date"`
	Specialty              string     `json:"specialty"`
	Notes                  *string    `json:"notes"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
}

func NewAppointment(userID, doctorName, hospitalMedicalRequest, appointmentDate, specialty, notes string) (*Appointment, error) {

	return &Appointment{
		ID:                     utils.NewID(),
		UserID:                 userID,
		DoctorName:             doctorName,
		HospitalMedicalRequest: hospitalMedicalRequest,
		AppointmentDate:        appointmentDate,
		Specialty:              specialty,
		Notes:                  &notes,
		CreatedAt:              time.Now(),
		UpdatedAt:              nil,
	}, nil
}
