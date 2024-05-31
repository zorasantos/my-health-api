package entity

import (
	"time"

	"github.com/zorasantos/my-health/utils"
)

type Examination struct {
	ID                     utils.ID   `json:"id"`
	UserID                 string     `json:"user_id"`
	Name                   string     `json:"name"`
	DoctorName             string     `json:"doctor_name"`
	HospitalMedicalRequest string     `json:"hospital_medical_request"`
	ExaminationDate        string     `json:"examination_date"`
	Notes                  string     `json:"notes"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
}

func NewExamination(userID, name, doctorName, hospitalMedicalRequest, examinationDate, notes string) (*Examination, error) {

	return &Examination{
		ID:                     utils.NewID(),
		UserID:                 userID,
		Name:                   name,
		DoctorName:             doctorName,
		HospitalMedicalRequest: hospitalMedicalRequest,
		ExaminationDate:        examinationDate,
		Notes:                  notes,
		CreatedAt:              time.Now(),
		UpdatedAt:              nil,
	}, nil
}
