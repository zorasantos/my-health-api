package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Group(func(r chi.Router) {
		// r.Post("/api/v1/login", handlers.Login)
		// r.Post("/api/v1/user", handlers.CreateUser)
	})

	return r
}
