package models

type User struct {
	UserID   int `gorm:"primary_key" json:"user_id"`
	Username int `gorm:"username;type:varchar(100)" json:"username"`
	Email    int `gorm:"email;type:varchar(200)" json:"email"`
	Password int `gorm:"password;type:varchar(255)" json:"password"`
	Movies []Movie
	Ratings []Rating
}
