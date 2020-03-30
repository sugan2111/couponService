package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	*mongo.Collection
}

// ConnectDB is a function to connect mongoDB
func ConnectDB() *mongo.Collection {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("myrestapi").Collection("coupons")

	return collection
}

func NewClient(uri string) MongoStore {
	return MongoStore{ConnectDB()}

}
