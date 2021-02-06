/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Exercise: To Lowercase
1. Look at the documentation of strings package
2. Find a function that changes letters to lowercase
3. Get the value from the command-line
4. Print the given value in lowercase letters

Output:
go run lowercase.go SHEPARD
shepard

Notes:

*/

func main() {
	input := strings.ToLower(os.Args[1])
	fmt.Println(input)
}