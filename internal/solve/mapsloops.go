/*
Exercise: Code Magnets
Produce the Output
The gold medal's rank is 1
The bronze medal's rank is 3
The silver medal's rank is 2
*/
package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "sliver": 2, "bronze": 3}

	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}
}