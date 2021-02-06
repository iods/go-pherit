/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
Exercise: Trim right
1. Look at the documentation for the strings package
2. Find a function that trims the spaces from only the right-most part of a given string
3. Trim it from the right part only
4. Print the number of characters it contains

Output:
3

Notes:

*/

func main() {
	name := "Rye         "
	name = strings.TrimRight(name, " ")
	lng := utf8.RuneCountInString(name)
	fmt.Println(lng)
}