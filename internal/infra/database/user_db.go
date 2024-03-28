package database

import (
	"database/sql"
	"log"

	"github.com/zorasantos/my-health/internal/entity"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	_, err := u.DB.Exec("INSERT INTO users (id, username, password, email, email_token, forgot_password_token, is_verified ) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.ID, user.Username, user.Password, user.Email, user.EmailToken, user.ForgotPasswordToken, user.IsVerified)

	if err != nil {
		return err
	}

	defer u.DB.Close()

	log.Println("User created successfully")

	return err
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	row := u.DB.QueryRow("SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.EmailToken, &user.ForgotPasswordToken, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)

	defer u.DB.Close()

	return &user, err
}
