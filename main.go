package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kjvmartinez/go-api-mongodb/config"
	"github.com/kjvmartinez/go-api-mongodb/database"
	"github.com/kjvmartinez/go-api-mongodb/handlers"
	"github.com/kjvmartinez/go-api-mongodb/middleware"
	"github.com/kjvmartinez/go-api-mongodb/services"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	database.InitDB(cfg.MongoDBURI)

	// Set JWT secret
	services.SetJWTSecret(cfg.JWTSecret)

	// Create router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORS())

	// Public routes (no auth required)
	router.POST("/api/auth/register", handlers.Register)
	router.POST("/api/auth/login", handlers.Login)

	// Protected routes (auth required)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/users", handlers.CreateUser)
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)
	}


	// Routes
	// router.POST("/api/users", handlers.CreateUser)
	// router.GET("/api/users", handlers.GetUsers)
	// router.GET("/api/users/:id", handlers.GetUser)
	// router.PUT("/api/users/:id", handlers.UpdateUser)
	// router.DELETE("/api/users/:id", handlers.DeleteUser)

	log.Println("🚀 Server running on http://localhost:" + cfg.Port)
	router.Run(":" + cfg.Port)
}