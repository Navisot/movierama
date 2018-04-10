package middlewares

import (
	"net/http"
	"github.com/navisot/movierama/app/config"
	"context"
)
type CustomMiddlewareHandler struct {}


func (m CustomMiddlewareHandler) AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		email := config.CheckSessionEmail(r)

		if email == "" {
			http.Redirect(w, r, "/", 302)
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}