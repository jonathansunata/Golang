package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LName string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
}
