package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zorasantos/my-health/config"

	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/internal/infra/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

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
	// r.SetTrustedProxies([]string{"187.58.71.4"})

	r.Post("/api/v1/login", handlers.Login)
	r.Post("/api/v1/user", handlers.CreateUser)

	http.ListenAndServe(":8080", r)

}
