package main

/*
Exercise:
Produce the following output:
0 19.07
1 0
2 25.2
0 a
1 b
2 c
*/

import "fmt"

func two() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	slice = append(slice, "x")
	slice = append(slice, "y", "z")

	for _, letter := range slice {
		fmt.Println(letter)
		// b
		// c
		// x
		// y
		// z
	}
}


func main() {
	// they are of different types, what do we do? float, int, and string.

	// first, init the var to hold the floats (0 can be a float, made empty when array is created)
	numbers := make([]float64, 3)
	numbers[0] = 19.07
	numbers[2] = 25.2

	for i, number := range numbers {
		fmt.Println(i, number)
		// 0 19.07
		// 1 0
		// 2 25.2
	}

	var letters = []string{"a", "b", "c"}

	for i, letter := range letters {
		fmt.Println(i, letter)
		// 0 a
		// 1 b
		// 2 c

	}

	two()
}