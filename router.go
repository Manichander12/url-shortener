package main

import "github.com/gorilla/mux"

func SetUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", ShortenUrlHandler).Methods("POST")
	r.HandleFunc("/{shortUrl}", RedirectHandler).Methods("GET")
	return r
}
