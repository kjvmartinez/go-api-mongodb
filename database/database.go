package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func InitDB(mongoDBURI string) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client
	log.Println("✅ Connected to MongoDB")
}

func GetCollection(dbName, collectionName string) *mongo.Collection {
	return DB.Database(dbName).Collection(collectionName)
}