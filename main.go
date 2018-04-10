package main

import (
	"github.com/navisot/movierama/app/config"
	"net/http"
	"github.com/navisot/movierama/app/database"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/navisot/movierama/app"
)

func main() {

	database.DB, database.Err = gorm.Open("postgres", "user=go_user password=go_password database=go_database sslmode=disable")

	if database.Err != nil {
		fmt.Println(database.Err.Error())
	} else {
		config.StartMigration()
	}

	router := app.NewRouter()

	http.ListenAndServe(":8024", router)
}
