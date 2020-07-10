package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathansunata/Golang/models"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{
		session: s,
	}
}

func (uc *UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	u := models.User{
		FName: "Jonathan",
		LName: "Sunata",
		Id:    1,
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

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", bs)
}
