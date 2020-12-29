/*
Exercise: Print the literals

1. Print a few integer literals
2. Print a few float literals
3. Print a few true and false bool literals
4. Print your name with a string literal
5. Print a non-english sentence with a string literal
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Integers:", 0, 1, 2, 3, 5, 8, 13) // print a few integer literals
	fmt.Println("Floats:", 3.14, 9.8, 143.0)       // print a few float literals
	fmt.Println("Booleans:", true, false)          // print a few true false literals
	fmt.Println("Your Name: Rye Miller!", )        // print your name with a string literal
	fmt.Println("Merhaba, adım  Rye Miller!")      // print a non-english (turkish) sentence with a string literal
}
