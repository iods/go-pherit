package main

import (
	"net/http"

	"github.com/iods/go-pherit/pkg/server"
)

func main() {
	mux := &servers.CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
