package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
errorCheck Function for reporting errors. */
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
viewHandler Handles the response for requests called to /guestbook. */
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("assets/view.html")
	errorCheck(err)
	err = html.Execute(writer, nil)
	errorCheck(err)
}

/*
main Starts the webserver and routes the requests to their appropriate function. */
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8150", nil)
	log.Fatal(err)
}