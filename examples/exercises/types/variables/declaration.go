/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Declare variables

1. Declare and print a variable with an int type
  * the variable's name should be: height
2. Declare and print a bool variable
  * the variable's name should be isOn
3. Declare and print a variable with a float64 type
  * the variable's name should be brightness
4. Declare and print a variable with a string type
  * print the type, and the string
5. Declare a variable
  * the variable's name should be isLiquid
  * discard it using a blank-identifier

Output:
1. 0
2. false
3. 0
4. s (string): ""
Notes:

*/

import "fmt"

func main() {
	var height int
	var isOn bool
	var brightness float64
	var s string
	var isLiquid bool

	fmt.Println(height)
	fmt.Println(isOn)
	fmt.Println(brightness)
	fmt.Printf("s (%T): %q\n", s, s)

	_ = isLiquid
}