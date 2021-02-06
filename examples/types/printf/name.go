/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import "fmt"

/*

Exercise:

1. Print your name and last name using Printf

Output:
My name is Rye and last name is Miller.

Notes:
  * store the formatting specifier (first arg of sprintf) in a variable

*/

func main() {
	message := "My first name is %s and last name is %s\n"
	fmt.Printf(message, "rye", "miller")
}
