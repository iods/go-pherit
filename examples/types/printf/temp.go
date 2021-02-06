/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import "fmt"

/*

Exercise:

1. Print the current temperature in your area using Printf

Output:
Temperature is 29.5 degrees.

Notes:
  * do not use the %v
  * output should NOT be like 29.0000000

*/

func main() {
	fmt.Printf("Temperature is %0.1f degrees.\n", 29.5)
}
