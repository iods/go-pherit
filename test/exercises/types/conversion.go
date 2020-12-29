/*
Exercise: Convert and Fix

1. Fix the code by using the conversion expression

Output:
15.5
10.5
9.5
1127

Notes:

*/
package main

import "fmt"

func main() {

	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	a = int(b)
	fmt.Println(float64(a) + b)

	age := 2
	fmt.Println(7.5 + float64(age))

	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))
}