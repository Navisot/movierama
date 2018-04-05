package config

import (
	"github.com/navisot/movierama/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func StartMigration(db *gorm.DB) {

	var allModels = make(map[string]interface{})
	allModels["User"] = &models.User{}
	allModels["Rating"] = &models.Rating{}
	allModels["Movie"] = &models.Movie{}

	for _, table := range allModels {
		db.DropTableIfExists(table)
		db.AutoMigrate(table)
	}



}
