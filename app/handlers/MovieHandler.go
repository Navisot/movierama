package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/navisot/movierama/app/config"
	"github.com/navisot/movierama/app/controllers"
	"github.com/navisot/movierama/app/models"
)

type MovieResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func SaveNewMovie(w http.ResponseWriter, req *http.Request) {

	email := req.Context().Value("email").(string)

	userID := controllers.GetUserIDByEmail(email)

	var movie models.Movie

	err := json.NewDecoder(req.Body).Decode(&movie)

	if err != nil {
		panic(err)
	}

	movie.UserID = userID

	_, err = controllers.SaveMovie(&movie)

	var Response MovieResponse

	if err != nil {
		panic(err)
	} else {
		Response.Status = "OK"
		Response.Code = 200
	}

	json.NewEncoder(w).Encode(&Response)

}

func GetAllMovies(w http.ResponseWriter, req *http.Request) {

	isLoggedIn := false

	// Must Know If User Is Logged In
	email := config.CheckSessionEmail(req)

	if email != "" {
		isLoggedIn = true
	}

	movies, err := controllers.GetAllMovies(isLoggedIn, email)

	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&movies)
}

func GetAllMoviesSorted(w http.ResponseWriter, req *http.Request) {

	p := mux.Vars(req)

	item := p["item"]

	isLoggedIn := false

	// Must Know If User Is Logged In
	email := config.CheckSessionEmail(req)

	if email != "" {
		isLoggedIn = true
	}

	movies, err := controllers.GetAllMoviesSorted(item, isLoggedIn, email)

	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&movies)
}

func GetAllMoviesTemplate(w http.ResponseWriter, r *http.Request) {

	is_logged_in := false
	username := ""

	email := config.CheckSessionEmail(r)

	if email != "" {
		is_logged_in = true
		username = controllers.GetUsernameByEmail(email)
	}
	cwd, _ := os.Getwd()
	p := controllers.MoviesTemplate{
		Title:    "HomePage",
		BasePath: "http://localhost:8024/",
		LoggedIn: is_logged_in,
		Username: username,
	}
	t, _ := template.ParseFiles(cwd + "/app/templates/movies.html")
	t.Execute(w, p)
}
