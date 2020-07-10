package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jonathansunata/Golang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{
		client: c,
	}
}

func (uc *UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	u := models.User{
		FName: "Jonathan",
		LName: "Sunata",
	}

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", bs)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	fmt.Println("%+v", u)
	collection := uc.client.Database("golangapideveloper").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}

func (uc *UserController) GetUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := uc.client.Database("golangapideveloper").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var u models.User

		err = cursor.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}
