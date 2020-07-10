package main

import (
	"net/http"

	"github.com/jonathansunata/Golang/controllers"
)

func main() {

	uc := controllers.NewUserController()

	http.HandleFunc("/user", uc.GetUser)
	http.ListenAndServe(":8080", nil)
}
