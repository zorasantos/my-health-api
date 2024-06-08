package database

import (
	"database/sql"
	"log"

	"github.com/zorasantos/my-health/internal/entity"
)

type Prescription struct {
	DB *sql.DB
}

func NewPrescription(db *sql.DB) *Prescription {
	return &Prescription{DB: db}
}

func (p *Prescription) Create(prescription *entity.Prescription) error {
	_, err := p.DB.Exec("INSERT INTO prescriptions (id, user_id, medicament_name, doctor_name, prescription_date, initial_date, final_date, notes, dosage, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", prescription.ID, prescription.UserID, prescription.MedicamentName, prescription.DoctorName, prescription.PrescriptionDate, prescription.InitialDate, prescription.FinalDate, prescription.Notes, prescription.Dosage, prescription.CreatedAt, prescription.UpdatedAt)

	if err != nil {
		return err
	}

	log.Println("Prescription created successfully")

	return err
}
