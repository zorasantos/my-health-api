package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zorasantos/my-health/db"
	"github.com/zorasantos/my-health/handlers"
	"github.com/zorasantos/my-health/middleware"
)

func main() {
	r := gin.Default()
	db.ConnectDB()
	// r.SetTrustedProxies([]string{"187.58.71.4"})

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
	}

	privateRoutes := r.Group("/protected")
	privateRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// privateRoutes.GET("/user", handlers.GetUser)
	}

	r.Run(":8080")

}
