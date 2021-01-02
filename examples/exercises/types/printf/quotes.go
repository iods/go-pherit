/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import "fmt"

/*

Exercise:

1.Print "hello world" with double quotes using printf

Output:
"hello world"

Notes:
  * output should be in quotes, not just hello world

*/

func main() {
	fmt.Printf("%q\n", "hello world")
}
