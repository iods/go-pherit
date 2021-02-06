/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Find the Perimeter of a Rectangle

1. Find the perimeter of a rectangle
  * it's width is 5
  * it's height is 6
2. Assign the result to the perimeter variable
3. Print the perimeter variable

Output:
22

Notes:
  * Formula = 2 times the width and height

*/

import "fmt"

func main() {
	var (
		perimeter int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)
}