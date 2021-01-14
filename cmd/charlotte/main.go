package main

import (
	"io"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World. It's Charlotte.")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
