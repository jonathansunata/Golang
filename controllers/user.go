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

	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	collection := uc.client.Database("golangapideveloper").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, u)
	json.NewEncoder(w).Encode(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", result)
}
