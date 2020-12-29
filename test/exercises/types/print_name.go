/*

Exercise:

1. Print your name and last name using Printf

Output:
My name is Rye and last name is Miller.

Notes:
  * store the formatting specifier (first arg of sprintf) in a variable

*/
package main

import "fmt"

func main() {
	message := "My first name is %s and last name is %s\n"
	fmt.Printf(message, "rye", "miller")
}
