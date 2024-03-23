package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zorasantos/my-health/config"
	"github.com/zorasantos/my-health/db"
	"github.com/zorasantos/my-health/handlers"
	"github.com/zorasantos/my-health/middleware"
)

func main() {
	r := gin.Default()

	_, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Failed to load config " + err.Error())
	}

	_, errorDB := db.ConnectDB()

	if errorDB != nil {
		log.Fatal("Failed to connect to database " + errorDB.Error())
	}

	if errorDB == nil {
		log.Println("Connected to database successfully")
	}
	// r.SetTrustedProxies([]string{"187.58.71.4"})

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/user", handlers.CreateUser)
	}

	privateRoutes := r.Group("/protected")
	privateRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// privateRoutes.GET("/user", handlers.GetUser)
	}

	r.Run(":8080")

}
