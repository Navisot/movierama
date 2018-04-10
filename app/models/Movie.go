package models

import "time"

type Movie struct {
	MovieID         int       `gorm:"primary_key" json:"movie_id"`
	Title           string    `gorm:"title;type:varchar(255)" json:"title"`
	Description     string    `gorm:"description" json:"description"`
	UserID          int       `json:"user_id"`
	Ratings         []Rating  `gorm:"foreignKey:MovieID"`
	User            User      `gorm:"foreignkey:UserID"`
	PublicationDate time.Time `gorm:"publication_date"`
}
