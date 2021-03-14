package errors

import (
	"fmt"
	"log"
	"os"
)

// Handle Function for reporting errors, checks if nil else panics.
func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// HandleV Handles an error, prints a custom message, and exits.
func HandleV(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}