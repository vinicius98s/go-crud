package database

import (
	"context"
	"fmt"
	"log"

	"go-crud/controllers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to the database
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to database: %v \n", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error checking the database connection: %v \n", err)
	}

	fmt.Println("Connected to database!")

	db := client.Database("go-crud")
	// Create collections by controllers
	controllers.UsersCollection(db)
}
