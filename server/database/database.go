package database

import (
	"context"
	"log"
	"os"
	"time"

	. "github.com/yokyann/PC3R_tetrio/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Query database
func List_Users() []User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var userList []User
	cursor, err := userCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//Iterating through the book elements
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatalln(err)
		}
		userList = append(userList, user)
	}

	return userList
}
func Find_User(pseudo string) *User {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *User
	filter := bson.D{{"pseudo", pseudo}}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}
func Create_User(user User) string {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userToPost := User{
		Id:    primitive.NewObjectID(),
		Pseudo:  user.Pseudo,
		Password: user.Password,
	}
	result, err := userCollection.InsertOne(ctx, userToPost)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}
