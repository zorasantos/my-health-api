package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zorasantos/my-health/config"

	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/internal/routes"
)

func main() {
	addr := ":8080"
	_, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Failed to load config " + err.Error())
	}

	_, errorDB := database.ConnectDB()

	if errorDB != nil {
		log.Fatal("Failed to connect to database " + errorDB.Error())
	}

	if errorDB == nil {
		log.Println("Connected to database successfully")
	}
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, routes.Router())

}
