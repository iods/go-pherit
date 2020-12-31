package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, Charlotte!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Bonjour, Charlotte!")
}

func germanHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Guten Tag, Charlotte!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/france", frenchHandler)
	http.HandleFunc("/german", germanHandler)
	err := http.ListenAndServe("localhost:8150", nil)
	log.Fatal(err)
}
