/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import (
	"fmt"
	"strings"
)

/*
Exercise: Trim It
1. Look at the documentation for the strings package
2. Find a function that trims a space from the given string
3. Trim the variable and print it out

Output:
The weather looks good.
I should go out and play.

Notes:

*/

func main() {
	msg := `

	The weather looks good.
I should go play outside.





`
	fmt.Println(strings.TrimSpace(msg))
}