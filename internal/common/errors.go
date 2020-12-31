package common

import "log"

/*
ErrorCheck Function for reporting errors. */
func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}