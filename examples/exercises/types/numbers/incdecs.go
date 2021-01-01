/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Incdecs
1. Increase the `counter` 5 times
2. Decrease the `factor` 2 times
3. Print the product of counter and factor

Output:
75

Notes:
  * use only incdec statements

*/
import "fmt"

func main() {
	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(float64(counter) * factor)
}