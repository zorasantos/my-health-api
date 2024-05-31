package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/zorasantos/my-health/config"

	"github.com/zorasantos/my-health/internal/infra/database"
	"github.com/zorasantos/my-health/internal/infra/handlers"
	authMiddleware "github.com/zorasantos/my-health/middleware"
)

func main() {
	addr := ":8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(httprate.LimitByIP(50, 1*time.Minute))

	_, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Failed to load config " + err.Error())
	}

	db, errorDB := database.ConnectDB()

	if errorDB != nil {
		log.Fatal("Failed to connect to database " + errorDB.Error())
	}

	if errorDB == nil {
		log.Println("Connected to database successfully")
	}
	userDB := database.NewUser(db)
	examinationDB := database.NewExamination(db)
	userHandler := handlers.NewUserHandler(userDB)
	examinationHandler := handlers.NewExaminationHandler(examinationDB)
	loginHandler := handlers.UserLoginHandler(userDB)

	r.Group(func(r chi.Router) {
		r.Post("/api/v1/login", loginHandler.Login)
		r.Post("/api/v1/user", userHandler.CreateUser)
		r.Get("/api/v1/user/{id}", userHandler.GetUser)
		r.Patch("/api/v1/user/{id}", userHandler.UpdateUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.AuthenticationMiddleware)
		r.Post("/api/v1/examination", examinationHandler.CreateExamination)
	})

	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, r)

}
