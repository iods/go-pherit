package main

import "fmt"

// simple defined type comedy error
type ComedyError string

// create a error method returning a string to satisfy the error interface
func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var err error // setup a var with the type (interface) of error
	err = ComedyError("This is an error.")
	fmt.Println(err)
}


