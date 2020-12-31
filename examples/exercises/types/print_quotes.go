/*

Exercise:

1.Print "hello world" with double quotes using printf

Output:
"hello world"

Notes:
  * output should be in quotes, not just hello world

*/
package main

import "fmt"

func main() {
	fmt.Printf("%q\n", "hello world")
}
