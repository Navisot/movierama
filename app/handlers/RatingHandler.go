package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	if success {
		json.NewEncoder(w).Encode(&RateMovieResponse{Message: "OK", Code: 200})
	} else {
		json.NewEncoder(w).Encode(&RateMovieResponse{Message: "FAIL", Code: 200})
	}

}
