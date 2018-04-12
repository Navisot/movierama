package models

type User struct {
	UserID   int      `gorm:"primary_key" json:"user_id"`
	Username string   `gorm:"username;type:varchar(100)" json:"username"`
	Email    string   `gorm:"email;type:varchar(200)" json:"email"`
	Password string   `gorm:"password;type:varchar(255)" json:"password"`
	Movies   []Movie  `gorm:"foreignKey:movie_id" json:"movie_id"`
	Ratings  []Rating `gorm:"foreignKey:rating_id" json:"rating_id"`
}
