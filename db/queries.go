package db

import (
	"errors"
	"log"

	"github.com/zorasantos/my-health/utils"
)

func CreateUser(username string, password string, email string) error {
	db, err := ConnectDB()
	userId, _ := utils.GenerateUUID()
	hashPassword, _ := utils.HashPassword(password)
	email_token, _ := utils.GenerateTokenEmail()
	is_verified := false
	forgot_password_token := ""

	if err != nil {
		return err
	}

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

	return nil
}
