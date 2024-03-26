package database

import (
	"errors"
	"log"

	"github.com/zorasantos/my-health/internal/entity"
)

func Create(user *entity.User) error {
	db, err := ConnectDB()

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (id, username, password, email, email_token, forgot_password_token, is_verified ) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.ID, user.Username, user.Password, user.Email, user.EmailToken, user.ForgotPasswordToken, user.IsVerified)

	if err != nil {
		return err
	}

	log.Println("User created successfully")

	defer db.Close()

	return err
}

func FindByEmail(email string) (entity.User, error) {
	db, err := ConnectDB()

	var user entity.User

	if err != nil {
		return user, errors.New("error connection db in get user")
	}

	row := db.QueryRow("SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.EmailToken, &user.ForgotPasswordToken, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, errors.New("User not found " + err.Error())
	}

	defer db.Close()

	return user, err
}
