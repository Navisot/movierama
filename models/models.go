package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	)

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
	Movies []Movie
	Ratings []Rating
}

type Movie struct {
	gorm.Model
	Title string
	Description string
	UserID int
	User User
	Ratings []Rating
}

type Rating struct {
	gorm.Model
	Like int
	Hate int
	MovieID int
	UserID int
	Movie Movie
	User User
}

