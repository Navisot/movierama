package main

import (
	"fmt"
	"github.com/navisot/movierama/app/config"
	"github.com/gorilla/mux"
	"github.com/navisot/movierama/app/handlers"
	"net/http"

	"github.com/navisot/movierama/app/database"
	"github.com/jinzhu/gorm"
)

func main() {

	database.DB, database.Err = gorm.Open("postgres", "user=go_user password=go_password database=go_database sslmode=disable")

	if database.Err != nil {
		fmt.Println(database.Err.Error())
	} else {
		config.StartMigration()
	}

	router := mux.NewRouter()

	router.HandleFunc("/user/register/{username}/{email}/{password}", handlers.NewUserRegistration).Methods("POST")

	http.ListenAndServe(":8084", router)
}
