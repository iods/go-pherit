package main

import (
	"fmt"
	"strings"
)

/*
paintNeeded Function to calculate the amount of paint (in
  liters) needed based on the area (meters) of a room.
*/
func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f Liters needed of paint.\n", area / 10)
}

/*
repeatLine Example function that takes two parameters, one a
  line (string) and two times (int), the output is then the
  line repeated x times.
*/
func repeatLine(line string, times int) {
	for i:= 0; i < times; i++ {
		fmt.Println(line)
	}
}

func main() {
	title := "Funk: Functions in Go."
	fmt.Println(strings.Title(title))

	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}