package common

import "log"

// HandleError Function for reporting errors, checks if nil else panics.
func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}