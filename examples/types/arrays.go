/*
Exercise:

1. Write down what you think the output will be.

Output:
42
0
108
c
a
b

Notes:

*/
package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
}
