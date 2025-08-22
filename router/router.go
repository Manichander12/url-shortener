package router

import (
	"github.com/gorilla/mux"
	"url-shortener/handler"
)

func SetUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", handler.ShortenUrlHandler).Methods("POST")
	r.HandleFunc("/{shortUrl}", handler.RedirectHandler).Methods("GET")
	return r
}
