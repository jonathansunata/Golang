package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathansunata/Golang/models"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
