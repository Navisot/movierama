package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetDatabaseInstance() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "user=go_user password=go_password database=go_database sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}
