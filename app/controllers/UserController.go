package controllers

import (
	"github.com/navisot/movierama/app/models"
	"github.com/jinzhu/gorm"
)

func RegisterUser(u *models.User) (*models.User, error) {

	db, err := gorm.Open("postgres", "user=go_user password=go_password database=go_database sslmode=disable")

	if err != nil {
		return nil,err
	}

	saveError := db.Debug().Save(u).Error


	if saveError != nil {
		return nil, saveError
	}

	return u, nil
}