package models

import "github.com/google/uuid"

type User struct {
	ID                  uuid.UUID `json:"id"`
	Username            string    `json:"username"`
	Password            string    `json:"password"`
	Email               string    `json:"email"`
	EmailToken          string    `json:"email_token"`
	ForgotPasswordToken string    `json:"forgot_password_token"`
	IsVerified          bool      `json:"is_verified"`
	CreatedAt           string    `json:"created_at"`
	UpdatedAt           *string   `json:"updated_at"`
}
