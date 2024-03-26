package entity

import (
	"time"

	"github.com/zorasantos/my-health/utils"
)

type User struct {
	ID                  utils.ID   `json:"id"`
	Username            string     `json:"username"`
	Password            string     `json:"-"`
	Email               string     `json:"email"`
	EmailToken          string     `json:"email_token"`
	ForgotPasswordToken *string    `json:"forgot_password_token"`
	IsVerified          bool       `json:"is_verified"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

func NewUser(username, password, email string) (*User, error) {
	hash, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	email_token, errorTokenEmail := utils.GenerateTokenEmail()

	if errorTokenEmail != nil {
		return nil, errorTokenEmail
	}

	return &User{
		ID:                  utils.NewID(),
		Username:            username,
		Password:            string(hash),
		Email:               email,
		EmailToken:          email_token,
		ForgotPasswordToken: nil,
		IsVerified:          false,
		CreatedAt:           time.Now(),
		UpdatedAt:           nil,
	}, nil
}
