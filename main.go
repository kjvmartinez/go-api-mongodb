package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kjvmartinez/go-api-mongodb/config"
	"github.com/kjvmartinez/go-api-mongodb/database"
	"github.com/kjvmartinez/go-api-mongodb/handlers"
	"github.com/kjvmartinez/go-api-mongodb/middleware"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	database.InitDB(cfg.MongoDBURI)

	// Create router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORS())

	// Routes
	router.POST("/api/users", handlers.CreateUser)
	router.GET("/api/users", handlers.GetUsers)
	router.GET("/api/users/:id", handlers.GetUser)
	router.PUT("/api/users/:id", handlers.UpdateUser)
	router.DELETE("/api/users/:id", handlers.DeleteUser)

	log.Println("🚀 Server running on http://localhost:" + cfg.Port)
	router.Run(":" + cfg.Port)
}