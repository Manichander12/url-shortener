package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	url_store map[string]string
	mutex = sync.RWMutex{}
)

type UrlRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:short_url`
}

func ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	var req UrlRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "invalid request url", http.StatusBadRequest)
		return																																																																			
	}

	shortKey := GenerateShortKey()

	mutex.Lock()						
	url_store[shortKey] = req.URL
	mutex.Unlock()

	resp := URLResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortKey),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func RedirectHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	shortKey := vars["shortUrl"]

	mutex.RLock()
	originalUrl, exists :=url_store[shortKey]
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}