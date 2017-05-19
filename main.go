package main

import (
	"net/http"
	"log"
	"regexp"
)

var githash string
var builddate string

func redirect(w http.ResponseWriter, r *http.Request) {
	var urlScheme string = "http"
	var host string = r.Host

	reg, err := regexp.Compile(`\:\d+$`)

	if err != nil {
		log.Printf("Error in regex parsing: %v", err)
	}

	if(r.Header.Get("X-Forwarded-Proto") == "https") {
		urlScheme = "https"
	}

	if(r.Header.Get("X-Forwarded-For") != "") {
		host = r.Header.Get("X-Forwarded-For")
	}

	safe := reg.ReplaceAllString(host, "")

	var fullUrl string = urlScheme + "://www." + safe + r.URL.Path

	log.Printf("Redirecting to: %v", fullUrl)
	http.Redirect(w, r, fullUrl, 301)
}

func main() {
	http.HandleFunc("/", redirect)

	log.Printf("Commit: %v, Built: %v", githash, builddate)
	log.Printf("Listening...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}