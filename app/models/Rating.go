package models

type Rating struct {
	RatingID int `gorm:"primary_key" json:"rating_id"`
	UserID int `gorm:"user_id" json:"user_id"`
	MovieID int `gorm:"movie_id" json:"user_id"`
	UpVote	int	`gorm:"upvote" json:"upvote"`
	DownVote	int	`gorm:"downvote" json:"downvote"`
	User User `gorm:"foreignkey:UserID"`
	Movie Movie `gorm:"foreignkey:MovieID"`
}
