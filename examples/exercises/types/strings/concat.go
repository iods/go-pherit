/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Raw Concatenation
1. Initialize the name variable by getting input from the command line.
2. Use the concatenation operator to concatenate it with the raw string literal

Output:
go run main.go Rye
hi Rye!
how are you?

Notes:

*/

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	msg := `hi ` + name + `
how are you?`
	fmt.Println(msg)
}