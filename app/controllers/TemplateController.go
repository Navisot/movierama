package controllers

type HomePageView struct {
	Title    string
	BasePath string
}

type LoginFormView struct {
	Title    string
	BasePath string
}

type RegisterFormView struct {
	Title    string
	BasePath string
}

type MoviesTemplate struct {
	Title    string
	BasePath string
	LoggedIn bool
	Username string
}
