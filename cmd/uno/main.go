/*
Go-Pherit: A Learning Project and SDK in Go
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/


/*
Package clause
Every go file can only belong to one package
 */
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

/*
Function References
Runnable, callable, reusable and executable blocks of code :)

Entrypoint that tells Go where to start executing
Allows go to execute the program

After compiling the code, Go runtime will first
  run this function you do not call the main
  function, Go calls it (__init__)
*/
func main () {
	// functions code gos here
	var name ="Gophers"
	fmt.Println("Hello, " + name)
	// in Go standard output is console
	// . after function accesses its functions
	fmt.Println("Rye Miller")
	fmt.Println("Tiffany Creamer")
}