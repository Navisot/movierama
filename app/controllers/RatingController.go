package controllers

import (
	"fmt"

	"github.com/navisot/movierama/app/database"
	"github.com/navisot/movierama/app/models"
)

func SaveNewRating(userEmail string, movieID int, vote int) bool {

	// Find UserID
	userID := GetUserIDByEmail(userEmail)

	// New User To Put Data
	userRates := models.Rating{}

	// Init counter
	var count = 0

	// Check If User Has Already Vote For This Movie
	database.DB.Debug().Model(&userRates).Where("user_id = ? and movie_id = ?", userID, movieID).Find(&userRates).Count(&count)

	if count >= 1 {
		// Update
		fmt.Println("UPDATE")
		userRates.Vote = vote
		database.DB.Model(&userRates).Where("user_id = ? and movie_id = ?", userID, movieID).Update(&userRates)
		return true
	} else {
		fmt.Println("INSERT")
		var Rating models.Rating
		Rating.MovieID = movieID
		Rating.UserID = userID
		Rating.Vote = vote

		err := database.DB.Model(&Rating).Save(&Rating).Error

		if err != nil {
			return false
		}

		return true

	}

}
