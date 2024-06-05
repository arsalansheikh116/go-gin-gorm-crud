package main

import (
	"github.com/gin-gonic/gin"

	"go-crud/config"
	"go-crud/middleware"

	// "go-crud/models"
	// "go-crud/middleware"
	"go-crud/routes"
)

func main() {
	// Initialize the database connection
	config.LoadEnv()
	config.InitializeDatabase()
	r := gin.Default()
	r.Use(middleware.JWTAuthMiddleware())

	// Routes
	routes.InitializeRoutes(r)

	r.Run(":8080")
}
