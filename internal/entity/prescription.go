package entity

import (
	"time"

	"github.com/zorasantos/my-health/utils"
)

type Prescription struct {
	ID               utils.ID   `json:"id"`
	UserID           string     `json:"user_id"`
	MedicamentName   string     `json:"medicament_name"`
	DoctorName       string     `json:"doctor_name"`
	PrescriptionDate string     `json:"prescription_date"`
	InitialDate      *string    `json:"initial_date"`
	FinalDate        *string    `json:"final_date"`
	Notes            *string    `json:"notes"`
	Dosage           string     `json:"dosage"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

func NewPrescription(userID, medicamentName, doctorName, prescriptionDate, initialDate, finalDate, notes, dosage string) (*Prescription, error) {

	return &Prescription{
		ID:               utils.NewID(),
		UserID:           userID,
		MedicamentName:   medicamentName,
		DoctorName:       doctorName,
		PrescriptionDate: prescriptionDate,
		InitialDate:      &initialDate,
		FinalDate:        &finalDate,
		Notes:            &notes,
		Dosage:           dosage,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}, nil
}
