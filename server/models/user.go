package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id   	primitive.ObjectID 		`json:"id"`
	Pseudo 	string 					`json:"pseudo"`
	Password 	string					`json:"password"`
}