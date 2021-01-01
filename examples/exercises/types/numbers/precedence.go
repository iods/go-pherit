/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Precedence
1. Change the expressions to produce the expected outputs.

Output:
20
-16
-25
0.5
24
15
40

Notes:
  * use parentheses to change the precedence

*/

import "fmt"

func main() {
	fmt.Println((10 + 5) - (5 - 10))
	fmt.Println(-10 + 0.5 - (1 + 5.5))
	fmt.Println(5 + 10 * (2 - 5))
	fmt.Println(0.5 * (2 - 1))

	fmt.Println(3 + 1 / 2 * 10 + 4) // 24 did not get this

	fmt.Println(10 / 2 * (10 % 7))
	fmt.Println(100 / (5. / 2))
}