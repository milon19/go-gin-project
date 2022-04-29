package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `json:"id" bson:"id"`
	UserId string             `json:"user_id" bson:"user_id"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
	Type   string             `json:"type" bson:"type"`
}
