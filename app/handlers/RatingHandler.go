package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/navisot/movierama/app/database"
	"github.com/navisot/movierama/app/models"

	"github.com/navisot/movierama/app/config"
	"github.com/navisot/movierama/app/controllers"

	"github.com/gorilla/mux"
)

type RateMovieResponse struct {
	Message string
	Code    int
}

func RateMovie(w http.ResponseWriter, r *http.Request) {

	success := false

	email := config.CheckSessionEmail(r)

	urlParams := mux.Vars(r)
	movieID, err := strconv.Atoi(urlParams["movie_id"])
	vote, err2 := strconv.Atoi(urlParams["vote"])

	if err != nil || err2 != nil {
		fmt.Println("ERROR TO CONVERT STRING TO INT")
	} else {
		success = controllers.SaveNewRating(email, movieID, vote)
	}

	userID := controllers.GetUserIDByEmail(email)

	result := &models.MovieResult{}

	// Get Movie Details To Pass In Response
	err3 := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = ? and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = ? and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id and mv.movie_id = ? GROUP BY u.user_id, mv.movie_id", movieID, movieID, userID, movieID).Scan(&result).Error

	if err3 != nil {
		fmt.Println("ERROR")
	}

	if success {
		json.NewEncoder(w).Encode(&result)
	} else {
		json.NewEncoder(w).Encode(&RateMovieResponse{Message: "FAIL", Code: 200})
	}

}

func MakeMovieWithoutVote(w http.ResponseWriter, r *http.Request) {

	success := false

	params := mux.Vars(r)

	movieID, err := strconv.Atoi(params["movie_id"])
	email := config.CheckSessionEmail(r)
	userID := controllers.GetUserIDByEmail(email)

	if err != nil {
		fmt.Println("ERROR CONVERT STRING TO INT")
	} else {
		success = controllers.SaveNewRating(email, movieID, 0)
	}

	result := &models.MovieResult{}

	// Get Movie Details To Pass In Response
	err3 := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = ? and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = ? and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id and mv.movie_id = ? GROUP BY u.user_id, mv.movie_id", movieID, movieID, userID, movieID).Scan(&result).Error

	if err3 != nil {
		fmt.Println("ERROR")
	}

	if success {
		json.NewEncoder(w).Encode(&result)
	} else {
		json.NewEncoder(w).Encode(&RateMovieResponse{Message: "FAIL", Code: 200})
	}
}
