package db

import (
	"errors"
	"log"
)

func CreateUser(username string, password string, email string) error {
	db, err := ConnectDB()

	if err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", username, password, email)

	if err != nil {
		logMsgErrorCreate := "failed to create user" + err.Error()
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
