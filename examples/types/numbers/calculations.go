/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Do some calculations.
1. Print the sum of 50 and 25
2. Print the difference of 50 and 15.5
3. Print the product of 50 and 0.5
4. Print the quotient of 50 and 0.5
5. Print the remainder of 25 and 3
6. Print the negation of 5 + 2

Output:
75
34.5
25
100
1

Notes:

*/
import "fmt"

func main() {
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))
}