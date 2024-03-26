package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
