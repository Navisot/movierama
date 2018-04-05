package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/navisot/movierama/app/models"
	"github.com/navisot/movierama/app/controllers"
	"fmt"
)


func NewUserRegistration(w http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)
	var NewUser = models.User{}
	json.NewDecoder(req.Body).Decode(&NewUser)
	NewUser.Username = params["username"]
	NewUser.Email = params["email"]
	NewUser.Password = params["password"]
	u,err := controllers.RegisterUser(&NewUser)

	if err != nil {
		fmt.Println("ERROR")
	}

	json.NewEncoder(w).Encode(u)


}
