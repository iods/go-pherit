/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Manipulate a Counter
1. Write the simplest line of code to increase the counter variable by 1.
2. Write the simplest line of code to decrease the counter variable by 1.
3. Write the simplest line of code to increase the counter variable by 5.
4. Write the simplest line of code to multiply the counter variable by 10.

Output:
10

Notes:

*/

import "fmt"

func main() {
	var counter int

	counter++     // 1
	counter--     // 2
	counter += 5  // 3
	counter *= 10 // 4
	counter /= 5  // 5

	fmt.Println(counter)
}