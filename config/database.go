package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = "db2025"
var MahasiswaCollection = "data_mahasiswa"
var UserCollection = "user"
var MongoString string

func init() {
	// Get the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting working directory:", err)
	}
	log.Println("Current working directory:", pwd)

	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		// Try loading from explicit path
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println("Error loading .env from parent directory:", err)
		}
	}

	MongoString = os.Getenv("MONGOSTRING")
	if MongoString == "" {
		log.Println("Warning: MONGOSTRING environment variable is not set")
	} else {
		log.Println("MONGOSTRING is set and has length:", len(MongoString))
	}
}

func MongoConnect(dbname string) (db *mongo.Database) {
	if MongoString == "" {
		fmt.Println("Error: MONGOSTRING is empty. Please check your .env file")
		return nil
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}

	// Ping the database to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Failed to ping MongoDB: %v\n", err)
		return nil
	}

	fmt.Println("Successfully connected to MongoDB")
	return client.Database(dbname)
}
