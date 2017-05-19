package main

import (
	"net/http"
	"log"
)

var githash string
var builddate string

func redirect(w http.ResponseWriter, r *http.Request) {
	var urlScheme string = "http"

	if(r.Header.Get("X-Forwarded-Proto") == "https") {
		urlScheme = "https"
	}

	var fullUrl string = urlScheme + "://www." + r.Host + r.URL.Path

	log.Printf("Redirecting to: %v", fullUrl)
	http.Redirect(w, r, fullUrl, 301)
}

func main() {
	http.HandleFunc("/", redirect)

	log.Printf("Commit: %v, Built: %v", githash, builddate)
	log.Printf("Listening...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}