/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import "fmt"

/*
Exercise: Windows Path
1. Update the program to use a raw string literal

Output:
c:\program files\duper super\fun.txt
c:\program files\really\funny.png

Notes:

*/

func main() {
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)
}