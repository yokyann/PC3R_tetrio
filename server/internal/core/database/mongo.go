package database

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Client, error) {
    // Set your MongoDB connection string
    uri := "your-mongodb-connection-string"

    // Set up a context and connect to MongoDB
    ctx := context.TODO()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    // Check if the connection was successful
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }
    log.Println("Connected to MongoDB!")

    return client, nil
}
