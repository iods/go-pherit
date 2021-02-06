/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/
package main

/*
Exercise: Simplify the assignments
1. Refactor the code to simplify the assignments.

Output:
3

Notes:
  * use only incdecs and assignment operators

*/

import "fmt"

func main() {
	width, height := 10, 2

	width += 1
	width += height
	width -= 1
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}