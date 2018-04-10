package models

import "time"

type MovieResult struct {
	UserID           int       `json:"user_id"`
	MovieID          int       `json:"movie_id"`
	Username         string    `json:"username"`
	UserEmail        string    `json:"user_email"`
	MovieTitle       string    `json:"movie_title"`
	MovieDescription string    `json:"movie_description"`
	MovieDate        time.Time `json:"movie_date"`
	MovieLikes       int       `json:"movie_likes"`
	MovieHates       int       `json:"movie_hates"`
	UserLike         int       `json:"user_like"`
}
