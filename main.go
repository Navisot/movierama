package main

import (
	"fmt"
	"github.com/navisot/movierama/app/config"
	"github.com/gorilla/mux"
	"github.com/navisot/movierama/app/handlers"
	"net/http"
)

func main() {

	db, err := config.GetDatabaseInstance()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		config.StartMigration(db)
	}

	router := mux.NewRouter()

	router.HandleFunc("/user/register/{username}/{email}/{password}", handlers.NewUserRegistration).Methods("POST")

	http.ListenAndServe(":8082", router)
}
