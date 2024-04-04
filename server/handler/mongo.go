// handler/mongo.go
package handler

import (
	"context"
	"log"
	"time"

	"github.com/yokyann/PC3R_tetrio/server/database"
	"github.com/yokyann/PC3R_tetrio/server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListUsers() []models.User {
	client := database.GetClient()
	userCollection := database.GetCollection(client, database.COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var userList []models.User
	cursor, err := userCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatalln(err)
		}
		userList = append(userList, user)
	}

	return userList
}

func FindUser(pseudo string) *models.User {
	client := database.GetClient()
	userCollection := database.GetCollection(client, database.COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User
	filter := bson.D{{"pseudo", pseudo}}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func CreateUser(user models.User) string {
	client := database.GetClient()
	userCollection := database.GetCollection(client, database.COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userToPost := models.User{
		Id:       primitive.NewObjectID(),
		Pseudo:   user.Pseudo,
		Password: user.Password,
	}
	result, err := userCollection.InsertOne(ctx, userToPost)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}
