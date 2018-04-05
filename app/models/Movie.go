package models

type Movie struct {
	MovieID int `gorm:"primary_key" json:"movie_id"`
	Title string `gorm:"title;type:varchar(255)" json:"title"`
	Description string  `gorm:"description" json:"description"`
	UserID int `json:"user_id"`
	User User	`gorm:"foreignkey:UserID"`
}
