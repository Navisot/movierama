package controllers

import (
	"fmt"

	"github.com/navisot/movierama/app/database"
	"github.com/navisot/movierama/app/helpers"
	"github.com/navisot/movierama/app/models"
)

func RegisterUser(u *models.User) (*models.User, error) {

	saveError := database.DB.Save(&u).Error

	if saveError != nil {
		return nil, saveError
	}

	return u, nil
}

func GetAllUsers() (*[]models.User, error) {

	users := []models.User{}

	database.DB.Find(&users)

	return &users, nil

}

func UserLogin(email, password string) bool {

	user := models.User{}

	err := database.DB.Model(&user).Where("email = ?", email).First(&user).Error

	if err != nil {
		return false
	}

	hashedPass := user.Password

	samePassword := helpers.ComparePasswords(hashedPass, []byte(password))

	if samePassword {
		return true
	}

	return false
}

func EmailOrUsernameExists(email, username string) bool {

	user := models.User{}

	var user_counter int

	database.DB.Model(&user).Where("email = ? or username = ?", email, username).Count(&user_counter)

	if user_counter == 1 {
		return true
	}

	return false
}

func GetUserIDByEmail(email string) int {

	user := models.User{}

	err := database.DB.Model(&user).Where("email = ?", email).Find(&user).Error

	if err != nil {
		fmt.Println("User Nor Found")
	}

	return user.UserID
}

func GetUsernameByEmail(email string) string {

	user := models.User{}

	err := database.DB.Model(&user).Where("email = ?", email).Find(&user).Error

	if err != nil {
		fmt.Println("User Nor Found")
	}

	return user.Username
}
