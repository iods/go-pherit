package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
indexHandler Function to manage the index route. http://localhost:8150/ */
func IndexHandler(response http.ResponseWriter, request *http.Request, _ httprouter.Params) { // empty params
	io.WriteString(response, "Soma homepage.")
}

func postHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Thanks, %s!\n", p.ByName("name"))
}





// localhost/ = dashboard of data
// localhost/add = add a symptom to track full form
// localhost/add/[symptom] = add a specific symptom to track POST
// localhost/update/[symptom] = update a symptom PUT
// localhost/delete/[symptom] = removes a symtom reference DELETE
// localhost/soma/[symptom] = GET a symptom reference


func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/soma/:name", postHandler)

	log.Fatal(http.ListenAndServe("localhost:8150", router))
}
