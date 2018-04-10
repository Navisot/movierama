package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/navisot/movierama/app/handlers"
	"github.com/navisot/movierama/app/middlewares"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Mode        string
	HandlerFunc http.HandlerFunc
	MiddleWares []mux.MiddlewareFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Get All Movies Template",
		"GET",
		"/",
		"WEB",
		handlers.GetAllMoviesTemplate,
		nil,
	},
	Route{
		"Get All Users",
		"GET",
		"/api/show/users",
		"API",
		handlers.GetAllUsers,
		nil,
	},
	Route{
		"Get All Movies",
		"GET",
		"/api/show/movies",
		"API",
		handlers.GetAllMovies,
		nil,
	},
	Route{
		"Get All Movies SORT BY item",
		"GET",
		"/api/show/movies/sort/{item}",
		"WEB",
		handlers.GetAllMoviesSorted,
		nil,
	},
	Route{
		"Save New Movie (title, description)",
		"POST",
		"/api/save/movie",
		"API",
		handlers.SaveNewMovie,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
	Route{
		"User Login",
		"POST",
		"/user/login",
		"WEB",
		handlers.LoginHandler,
		nil,
	},
	Route{
		"User Login Form",
		"GET",
		"/user/login",
		"WEB",
		handlers.LoginFormHandler,
		nil,
	},
	Route{
		"User Logout",
		"GET",
		"/user/logout",
		"WEB",
		handlers.LogoutHandler,
		nil,
	},
	Route{
		"User Registration Form",
		"GET",
		"/user/register",
		"WEB",
		handlers.NewUserRegistrationForm,
		nil,
	},
	Route{
		"User Registration WEB",
		"POST",
		"/user/register",
		"WEB",
		handlers.WebNewUserRegistration,
		nil,
	},
}
