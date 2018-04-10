package config

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var SessionName = "movierama_session"

var cookieGenerator = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func SetSession(email string, response http.ResponseWriter) {

	value := map[string]string{
		"email": email,
	}
	encodedCookie, err := cookieGenerator.Encode(SessionName, value)

	if err == nil {
		cookie := &http.Cookie{
			Name:  SessionName,
			Value: encodedCookie,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   SessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}


func CheckSessionEmail(request *http.Request) (email string) {

	// Get Cookie, Error
	cookie, err := request.Cookie(SessionName)

	// Initialize Cookie Value To Get From Cookie Decode Method
	cookieValue := make(map[string]string)

	if err == nil {
		err = cookieGenerator.Decode(SessionName,cookie.Value, &cookieValue)
		if err == nil {
			email = cookieValue["email"]
		}
	}

	return email
}