package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection
var COLLECTION = "Users"

func GetClient() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	//getting context
	if client != nil {
		return client
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//getting client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectioName string) *mongo.Collection {
	if userCollection != nil {
		return userCollection
	}
	bookCollection := client.Database("tetrio_opgg").Collection(collectioName)
	return bookCollection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
