package entity

import (
	"time"

	"github.com/zorasantos/my-health/utils"
)

type User struct {
	ID                  utils.ID `json:"id"`
	Username            string   `json:"username"`
	Password            string   `json:"-"`
	Email               string   `json:"email"`
	EmailToken          string   `json:"email_token"`
	ForgotPasswordToken *string  `json:"forgot_password_token"`
	IsVerified          bool     `json:"is_verified"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           *string  `json:"updated_at"`
}

func NewUser(username, password, email string) (*User, error) {
	hash, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	email_token, errorEmail := utils.GenerateTokenEmail()

	if errorEmail != nil {
		return nil, errorEmail
	}

	return &User{
		ID:                  utils.NewID(),
		Username:            username,
		Password:            string(hash),
		Email:               email,
		EmailToken:          email_token,
		ForgotPasswordToken: nil,
		IsVerified:          false,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:           nil,
	}, nil
}
