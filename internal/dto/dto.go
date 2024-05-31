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

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
}
