package error

import (
	"fmt"
	"log"
	"os"
)

// HandleError Function for reporting errors, checks if nil else panics.
func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Handle Reports and handles errors.
func Handle(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}