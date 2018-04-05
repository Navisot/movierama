package config

import (
	"github.com/navisot/movierama/app/models"
	_ "github.com/lib/pq"
)

func StartMigration(db *Database) {

	var allModels = make(map[string]interface{})
	allModels["User"] = &models.User{}
	allModels["Rating"] = &models.Rating{}
	allModels["Movie"] = &models.Movie{}

	for _, table := range allModels {
		db.DB.DropTableIfExists(table)
		db.DB.AutoMigrate(table)
	}



}
