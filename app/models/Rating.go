package models

type Rating struct {
	RatingID int   `gorm:"primary_key" json:"rating_id"`
	UserID   int   `gorm:"user_id" json:"user_id"`
	MovieID  int   `gorm:"movie_id" json:"user_id"`
	Vote     int   `gorm:"vote" json:"vote"`
	User     User  `gorm:"foreignkey:UserID"`
	Movie    Movie `gorm:"foreignkey:MovieID"`
}
