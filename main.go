package main

import (
	"fmt"
	"github.com/navisot/movierama/app/config"
)

func main() {

	db, err := config.GetDatabaseInstance()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		config.StartMigration(db)
	}
}
