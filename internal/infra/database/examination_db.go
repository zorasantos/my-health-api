package database

import (
	"database/sql"
	"log"

	"github.com/zorasantos/my-health/internal/entity"
)

type Examination struct {
	DB *sql.DB
}

func NewExamination(db *sql.DB) *Examination {
	return &Examination{DB: db}
}

func (e *Examination) Create(examination *entity.Examination) error {
	_, err := e.DB.Exec("INSERT INTO examinations (id, user_id, name, doctor_name, hospital_medical_request, examination_date, notes) VALUES ($1, $2, $3, $4, $5, $6, $7)", examination.ID, examination.UserID, examination.Name, examination.DoctorName, examination.HospitalMedicalRequest, examination.ExaminationDate, examination.Notes)

	if err != nil {
		return err
	}

	log.Println("Examination created successfully")

	return err
}
