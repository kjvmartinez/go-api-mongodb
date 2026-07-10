package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI   string
	DatabaseName string
	Port         string
	JWTSecret    string
}

func LoadConfig() *Config {
	godotenv.Load()

	return &Config{
		MongoDBURI:   os.Getenv("MONGODB_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}
}