package db

import (
	"errors"
	"log"

	"github.com/zorasantos/my-health/models"
	"github.com/zorasantos/my-health/utils"
)

func CreateUser(username string, password string, email string) error {
	db, err := ConnectDB()

	if err != nil {
		return errors.New("error connection db create user " + err.Error())
	}

	userId, errorUUID := utils.GenerateUUID()

	if errorUUID != nil {
		log.Println(errorUUID)
		return errors.New("failed to generate uuid in create user " + errorUUID.Error())
	}

	hashPassword, errorHash := utils.HashPassword(password)

	if errorHash != nil {
		log.Println(errorHash)
		return errors.New("failed to hash password " + errorHash.Error())
	}

	email_token, errorEmail := utils.GenerateTokenEmail()

	if errorEmail != nil {
		log.Println(errorEmail)
		return errors.New("failed to generate email token " + errorEmail.Error())

	}

	var is_verified bool
	var forgot_password_token string

	result, err := db.Exec("INSERT INTO users (id, username, password, email, email_token, forgot_password_token, is_verified ) VALUES ($1, $2, $3, $4, $5, $6, $7)", userId, username, hashPassword, email, email_token, forgot_password_token, is_verified)

	if err != nil {
		logMsgErrorCreate := err.Error()
		log.Println(logMsgErrorCreate)
		return errors.New(logMsgErrorCreate)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		logMessage := "failed to get affected rows"
		log.Println(logMessage)
		return errors.New(logMessage)
	}

	if affectedRows == 0 {
		logMessage := "failed to create user"
		log.Println(logMessage)
		return errors.New(logMessage)
	}

	log.Println("User created successfully")

	defer db.Close()

	return err
}

func GetUser(email string) (models.User, error) {
	db, err := ConnectDB()

	var user models.User

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
