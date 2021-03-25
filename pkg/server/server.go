package server

import (
	"log"
	"net/http"

	"github.com/iods/go-pherit/internal/env"
	"github.com/iods/go-pherit/internal/handles"

)

// Run Function to initiate the web server and start the application.
func Run() {
	http.HandleFunc("/", handles.IndexPage)
	http.HandleFunc("/about", handles.AboutPage)
	log.Fatal(http.ListenAndServe(env.Port(), nil))
}