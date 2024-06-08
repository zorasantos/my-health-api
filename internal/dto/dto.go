package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateExaminationDTO struct {
	Name                   string `json:"name"`
	DoctorName             string `json:"doctor_name"`
	HospitalMedicalRequest string `json:"hospital_medical_request"`
	ExaminationDate        string `json:"examination_date"`
	Notes                  string `json:"notes"`
}

type CreateAppointmentDTO struct {
	DoctorName             string `json:"doctor_name"`
	HospitalMedicalRequest string `json:"hospital_medical_request"`
	AppointmentDate        string `json:"appointment_date"`
	Specialty              string `json:"specialty"`
	Notes                  string `json:"notes"`
}
type CreatePrescriptionDTO struct {
	MedicamentName   string `json:"medicament_name"`
	DoctorName       string `json:"doctor_name"`
	PrescriptionDate string `json:"prescription_date"`
	InitialDate      string `json:"initial_date"`
	FinalDate        string `json:"final_date"`
	Notes            string `json:"notes"`
	Dosage           string `json:"dosage"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
}
