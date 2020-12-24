package main

import (
	"fmt"
	"log"
	"strings"
)

/*
paintNeeded Function to calculate the amount of paint (in
  liters) needed based on the area (meters) of a room.
*/
func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
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

func (n *Number) Doubles() {
	*n *= 2
}

func main() {
	title := "Funk: Functions in Go."
	fmt.Println(strings.Title(title))

	amounts := Number(4)
	amounts.Double()


	fmt.Println(amounts)

	var amount, total float64

	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	} else {

	}
	fmt.Printf("%.2f liters needed.\n", amount)
	total += amount
}