package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/navisot/movierama/app/database"

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

	var Response MovieResponse

	email := req.Context().Value("email").(string)

	userID := controllers.GetUserIDByEmail(email)

	var movie models.Movie

	err := json.NewDecoder(req.Body).Decode(&movie)

	// Check If Movie Has Description N Title
	if movie.Title == "" || movie.Description == "" {
		Response.Status = "VALIDATOR_ERROR"
		Response.Code = 200
		json.NewEncoder(w).Encode(&Response)
		return
	}

	if err != nil {
		panic(err)
	}

	// Save Movie
	movie.UserID = userID
	movie.PublicationDate = time.Now()

	_, err = controllers.SaveMovie(&movie)

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

type MoviesAPIResponse struct {
	Movies         []models.MovieResult
	SpecificUserID int
	SelectedUser   string
}

func GetAllMoviesSorted(w http.ResponseWriter, req *http.Request) {

	response := MoviesAPIResponse{}

	p := mux.Vars(req)

	item := p["item"]

	userID, _ := strconv.Atoi(p["user_id"])

	isLoggedIn := false

	// Must Know If User Is Logged In
	email := config.CheckSessionEmail(req)

	if email != "" {
		isLoggedIn = true
	}

	movies, searchForUser, err := controllers.GetAllMoviesSorted(item, isLoggedIn, email, userID)

	if err != nil {
		log.Println(err)
	}

	if searchForUser {
		u := models.User{}
		database.DB.Where("user_id = ?", userID).First(&u)
		response.Movies = movies
		response.SpecificUserID = userID
		response.SelectedUser = u.Username
	} else {
		response.Movies = movies
		response.SpecificUserID = 0
	}

	json.NewEncoder(w).Encode(&response)
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
