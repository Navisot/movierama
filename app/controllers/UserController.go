package controllers

import (
	"github.com/navisot/movierama/app/models"
	"github.com/navisot/movierama/app/database"
)

func RegisterUser(u *models.User) (*models.User, error) {

	saveError := database.DB.Debug().Save(u).Error

	if saveError != nil {
		return nil, saveError
	}

	return u, nil
}