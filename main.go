package main

import (
	"log"
	"net/http"
	"url-shortener/router"
)

func main() {
	router := router.SetUpRouter()
	log.Println("Server started at : 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
