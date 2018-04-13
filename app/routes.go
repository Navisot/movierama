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
	HandlerFunc http.HandlerFunc
	MiddleWares []mux.MiddlewareFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Get All Movies Template",
		"GET",
		"/",
		handlers.GetAllMoviesTemplate,
		nil,
	},
	Route{
		"Rate Movie",
		"POST",
		"/api/rate/movie/{movie_id}/{vote}",
		handlers.RateMovie,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
	Route{
		"Get All Users",
		"GET",
		"/api/show/users",
		handlers.GetAllUsers,
		nil,
	},
	Route{
		"Get User Status",
		"GET",
		"/api/user/status",
		handlers.GetUserStatus,
		nil,
	},
	Route{
		"Get All Movies",
		"GET",
		"/api/show/movies",
		handlers.GetAllMovies,
		nil,
	},
	Route{
		"Get All Movies SORT BY item",
		"GET",
		"/api/show/movies/sort/{item}/{user_id}",
		handlers.GetAllMoviesSorted,
		nil,
	},
	Route{
		"Save New Movie (title, description)",
		"POST",
		"/api/save/movie",
		handlers.SaveNewMovie,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
	Route{
		"User Login",
		"POST",
		"/user/login",
		handlers.LoginHandler,
		nil,
	},
	Route{
		"User Login Form",
		"GET",
		"/user/login",
		handlers.LoginFormHandler,
		nil,
	},
	Route{
		"User Logout",
		"GET",
		"/user/logout",
		handlers.LogoutHandler,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
	Route{
		"User Registration Form",
		"GET",
		"/user/register",
		handlers.NewUserRegistrationForm,
		nil,
	},
	Route{
		"User Registration WEB",
		"POST",
		"/user/register",
		handlers.WebNewUserRegistration,
		nil,
	},
	Route{
		"User Registration WEB",
		"POST",
		"/api/unhate/movie/{movie_id}",
		handlers.MakeMovieWithoutVote,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
	Route{
		"User Registration WEB",
		"POST",
		"/api/unlike/movie/{movie_id}",
		handlers.MakeMovieWithoutVote,
		[]mux.MiddlewareFunc{middlewares.CustomMiddlewareHandler{}.AuthenticateMiddleware},
	},
}
