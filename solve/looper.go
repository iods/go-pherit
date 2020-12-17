package main

import "fmt"

func main() {
	// until x is less than or equal to 3, print
	for x :=1; x <= 3; x++ {
		fmt.Println(x) // 123
	}

	// until x is less than or equal to 3, print
	for x := 2; x <= 3; x++ {
		fmt.Println(x) // 23
	}

	// until x is = 6, add 2
	for x := 1; x <=6; x+= 2 {
		fmt.Println(x)
	}

	// starting from 3, until x is greater or equal to 1, subtract 1
	for x := 3; x >= 1; x-- {
		fmt.Println(x) // 321
	}

	// starting from 1, up to 3, add 1
	for x := 1; x < 3; x++ {
		fmt.Println(x) // 12
	}

	// starting at 1, until greater or equal to three, add one
	for x := 1; x >= 3; x++ {
		fmt.Println(x) // doesnt print nuthin
	}
}
