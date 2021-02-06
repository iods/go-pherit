/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise:

1. Change `color` variable's value to "blue"
2. Change value of `color` variable from "blue" to "red"
  * do not assign "red" to `color` directly (not color = "red")
3. Multiply 3.14 with 2 and assign it to `n` variable

Output:
1. blue
2. red
3. 6.28

Notes:

*/

import "fmt"

func main() {
	color := "green"
	color = "blue"

	color = "r" + "e" + "d"

	n := 0.
	n = 3.14 * 2

	fmt.Println(color)
	fmt.Println(n)
}