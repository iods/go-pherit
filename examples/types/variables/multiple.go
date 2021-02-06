/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Declare variables

1. Declare and print two vars using multiple variable declaration
  * the first variable's name should be: active
  * the second variable's name should be: delta
2. Declare them a bool and an int variable
3. Declare and initialize two string variables using multiple variable declaration
  * use the type once (single line) while declaring the variables
  * the first variable's name should be firstName
  * the second variable's name should be lastName

Output:
1. false 0
2. "" ""

Notes:

*/

import "fmt"

func main() {
	var (
		active bool
		delta int
	)
	var firstName, lastName string
	fmt.Println(active, delta)
	fmt.Printf("%q %q\n", firstName, lastName)
}