package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
	Password string           `bson:"password" json:"-"`
	Age   int                `bson:"age" json:"age"`
}

// Request models (for API input)
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"min=0,max=150"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}