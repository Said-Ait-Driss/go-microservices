package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		uri = "mongodb://localhost:27017/?directConnection=true"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("fatal load mongo db url ")
	}

	fmt.Println("connect to mongo db successfully")

	return client
}

var Client *mongo.Client = DBInstance()
