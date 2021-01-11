package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

/*
indexHandler Function to manage the index route. http://localhost:8150/ */
func IndexHandler(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	write(response, "Charlotte homepage.")
}

// englishHandler
func englishHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	write(writer, "Hello, Charlotte!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	write(writer, "Bonjour, Charlotte!")
}

func germanHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	write(writer, "Guten Tag, Charlotte!")
}

func postHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Thanks, %s!\n", p.ByName("name"))
}

func main() {

	r := httprouter.New()

	r.GET("/", IndexHandler)
	r.GET("/english", englishHandler)
	r.GET("/french", frenchHandler)
	r.GET("/german", germanHandler)
	r.GET("/charlotte/:name", postHandler)

	err := http.ListenAndServe("localhost:9111", r)
	log.Fatal(err)
}
