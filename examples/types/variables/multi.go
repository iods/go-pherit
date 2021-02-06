/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise:

1. Assign "go" to lang variable and assign 2 to version variable
  * use multiple assignment statement
2. Assign the correct values to the variables to match the output
3. Declare two vars using multiple short declaration assignment
  * init the vars using a multi() function
  * discard the first vars value in the declaration
  * print the second variable
4. Change color to "orange" and color2 to "green" at the same time
  * using multiple assignment
5. Swap the values of `red` and `blue` variables

Output:
go version 2
Air is good on Mars
It's true
It is 19.5 degrees
4
orange green
blue red

Notes:

*/

import "fmt"

func multi() (int, int) {
	return 5, 4
}

func main() {
	var (
		isTrue bool
		lang string
		planet string
		temp float64
		version int
	)

	_, b := multi()

	color, color2 := "red", "blue"
	color, color2 = "orange", "green"

	red, blue := "red", "blue"
	red, blue = blue, red

	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	isTrue, planet, temp = true, "Mars", 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	fmt.Println(b)

	fmt.Println(color, color2)

	fmt.Println(red, blue)
}