package database

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func InitDB(mongoDBURI string) {

	// Create TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create client options
	clientOpts := options.Client().
		ApplyURI(mongoDBURI).
		SetTLSConfig(tlsConfig).
		SetServerSelectionTimeout(30 * time.Second). // Increase timeout to 30 seconds
		SetConnectTimeout(30 * time.Second)

	// clientOpts := options.Client().
	// 	ApplyURI(mongoDBURI).
	// 	SetServerSelectionTimeout(10 * time.Second) // Increase timeout

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	DB = client
	log.Println("✅ Connected to MongoDB")
}

func GetCollection(dbName, collectionName string) *mongo.Collection {
	return DB.Database(dbName).Collection(collectionName)
}