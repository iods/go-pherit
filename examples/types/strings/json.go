/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Windows Path
1. Update the program to use a raw string literal

Output:

Notes:

*/

import "fmt"

func main() {
	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	fmt.Println(json)

	json = `
{
	"Items": [{
		"Items": {
			"name": "Teddy Bear"
		}
	}]
}
`
	fmt.Println(json)
}