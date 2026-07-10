package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/kjvmartinez/go-api-mongodb/database"
	"github.com/kjvmartinez/go-api-mongodb/models"
	"github.com/kjvmartinez/go-api-mongodb/services"
)

// const dbName = "go_api_db"
// const collectionName = "users"

// Register user
func Register(c *gin.Context) {
	var req models.RegisterRequest

	// Bind and validate input
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Validate struct
	if err := services.ValidateStruct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection(dbName, collectionName)

	// Check if email already exists
	count, err := collection.CountDocuments(context.Background(), bson.M{"email": req.Email})
	if err != nil || count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error processing password"})
		return
	}

	// Create user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      req.Age,
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate token
	userID := result.InsertedID.(primitive.ObjectID)
	token, err := services.GenerateToken(userID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	user.ID = userID
	c.JSON(http.StatusCreated, models.LoginResponse{
		Token: token,
		User:  user,
	})
}

// Login user
func Login(c *gin.Context) {
	var req models.LoginRequest

	// Bind and validate input
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Validate struct
	if err := services.ValidateStruct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection(dbName, collectionName)

	// Find user by email
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Verify password
	if !services.VerifyPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate token
	token, err := services.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
		User:  user,
	})
}