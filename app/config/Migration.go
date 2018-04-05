package config

import (
	"github.com/navisot/movierama/app/models"
	_ "github.com/lib/pq"
	"github.com/navisot/movierama/app/database"
)

func StartMigration() {

	var allModels = make(map[string]interface{})
	allModels["User"] = &models.User{}
	allModels["Rating"] = &models.Rating{}
	allModels["Movie"] = &models.Movie{}

	for _, table := range allModels {
		database.DB.DropTableIfExists(table)
		database.DB.AutoMigrate(table)
	}



}
