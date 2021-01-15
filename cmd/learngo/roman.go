package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/iods/go-pherit/internal/roman"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Split(r.URL.Path, "/") // split the URL by the slash
		if url[1] == "roman_number" { // if request is GET w/ correct syntax
			number, _ := strconv.Atoi(strings.TrimSpace(url[2]))
			if number == 0 || number > 10 {
				// return a Not Found status if not in the array
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found buddy."))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(roman.Numerals[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request buddy."))
		}
	})

	server := &http.Server{
		Addr: ":8000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}