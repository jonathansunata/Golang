package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id"`
	FName string             `json:"first_name"`
	LName string             `json:"last_name"`
}
