package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/navisot/movierama/app/config"
	"github.com/navisot/movierama/app/controllers"
	"github.com/navisot/movierama/app/helpers"
	"github.com/navisot/movierama/app/models"
)

func WebNewUserRegistration(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	validRecord := UserValidation(email, password, username)

	if validRecord == false {
		http.Redirect(w, r, "/user/register", 302)
		return
	}

	us := models.User{
		Username: username,
		Email:    email,
		Password: helpers.HashAndSalt([]byte(password)),
	}

	_, err := controllers.RegisterUser(&us)

	if err != nil {
		fmt.Println("ERROR")
	}

	http.Redirect(w, r, "/", 302)

}

func GetAllUsers(w http.ResponseWriter, req *http.Request) {

	users, err := controllers.GetAllUsers()

	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&users)
}

func HomePageTemplate(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := controllers.HomePageView{Title: "HomePage", BasePath: "http://localhost:8024/"}
	t, _ := template.ParseFiles(cwd + "/app/templates/base.html")
	t.Execute(w, p)
}

func NewUserRegistrationForm(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := controllers.RegisterFormView{Title: "Register User", BasePath: "http://localhost:8024/"}
	t, _ := template.ParseFiles(cwd + "/app/templates/register.html")
	t.Execute(w, p)
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := controllers.LoginFormView{Title: "User Login", BasePath: "http://localhost:8024/"}
	t, _ := template.ParseFiles(cwd + "/app/templates/login.html")
	t.Execute(w, p)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	redirectTarget := "/"

	if helpers.ValidEmail(email) == false {
		http.Redirect(w, r, "/user/login", 302)
		return
	}

	findUser := controllers.UserLogin(email, password)

	if findUser {
		config.SetSession(email, w)
		redirectTarget = "/"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	config.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}

func UserValidation(email, password, username string) bool {

	password_length := len(password)

	valid_email := helpers.ValidEmail(email)

	exists := controllers.EmailOrUsernameExists(email, username)

	correctPassword := helpers.Between2Numbers(password_length, 5, 15)

	if exists == true || correctPassword == false || valid_email == false {
		return false
	}

	return true
}
