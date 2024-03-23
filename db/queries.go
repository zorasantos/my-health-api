package db

import (
	"errors"
	"log"

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

	is_verified := false
	forgot_password_token := ""

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
