package main

import (
	"fmt"
	"github.com/iods/go-pherit/internal/common"
	"github.com/iods/go-pherit/internal/structs"
	"github.com/iods/go-pherit/pkg/datafile"
	"html/template"
	"log"
	"net/http"
	"os"
)

/*
viewHandler Handles the response for requests called to /guestbook. */
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := datafile.GetStrings("signatures.txt")
	html, err := template.ParseFiles("assets/view.html")
	errors.ErrorCheck(err)

	guestbook := structs.Guestbook{
		SignatureCount: len(signatures),
		Signatures: signatures,
	}
	err = html.Execute(writer, guestbook)
	errors.ErrorCheck(err)
}

/*
newHandler Routes a request to the form for accepting new signatures. */
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("assets/new.html")
	errors.ErrorCheck(err)

	err = html.Execute(writer, nil)
	errors.ErrorCheck(err)
}

/*
createHandler Assigns the route /guestbook/create to a POST function. */
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("Signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	errors.ErrorCheck(err)
	_, err = fmt.Fprintln(file, signature)
	errors.ErrorCheck(err)
	err = file.Close()
	errors.ErrorCheck(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

/*
main Starts the webserver and routes the requests to their appropriate function. */
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8150", nil)
	log.Fatal(err)
}