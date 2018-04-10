package app

import (
	"github.com/gorilla/mux"
	"flag"
	"net/http"
)

func NewRouter() *mux.Router {

	var dir string

	flag.StringVar(&dir, "dir", "static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc

		if route.MiddleWares != nil {
			for m := range route.MiddleWares {
				handler = route.MiddleWares[m](handler)
			}
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	return router
}