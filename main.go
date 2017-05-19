package main

import (
	"net/http"
	"net/http/httputil"
	"log"
	"regexp"
)

var githash string
var builddate string

func redirect(w http.ResponseWriter, r *http.Request) {
	// debug the request
	requestDump, _ := httputil.DumpRequest(r, true)

	log.Println(string(requestDump))

	var urlScheme string = "http"
	var host string = r.Host

	reg, err := regexp.Compile(`\:\d+$`)

	if err != nil {
		log.Printf("Error in regex parsing: %v", err)
	}

	if(r.Header.Get("X-Forwarded-Proto") == "https") {
		urlScheme = "https"
	}

	if(r.Header.Get("X-Forwarded-Host") != "") {
		host = r.Header.Get("X-Forwarded-Host")
	}

	safe := reg.ReplaceAllString(host, "")

	var fullUrl string = urlScheme + "://www." + safe + r.URL.Path

	log.Printf("Redirecting to: %v", fullUrl)
	http.Redirect(w, r, fullUrl, 301)
}

// health endpoint
func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to: /healthz")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{ \"health\": \"ok\" }"))
}

func main() {
	http.HandleFunc("/healthz", health)
	http.HandleFunc("/", redirect)

	log.Printf("Commit: %v, Built: %v", githash, builddate)
	log.Printf("Listening...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}