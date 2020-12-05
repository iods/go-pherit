// Dark Society - Go Development

// Package clause
// Every go file can only belong to one package
// This allows go to make it executable
package main

/*
Import Statement/Declaration
Import statement makes more packages available
  for this .go file/program.

Import "fmt" lets you access fmt package's functionality
  in this .go file. It is short for formatting and
  is part of the standard library.
 */
import "fmt"

// function - reusable and executable block of code :)
// entrypoint that tells Go where to start executing
// allows go to execute the program
// After compiling the code,
// Go runtime will first run this function
// you do not call the main function, Go calls it (__init__)
func main () {
	// functions code gos here
	fmt.Println("Hello, Gophers")
	// in Go standard output is console
	// . after function accesses its functions

	/* Exercises */
	fmt.Println("Rye Miller")
	fmt.Println("Tiffany Creamer")
}

