/*
Golph Calculates the average of several numbers.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:] // get slice of strings, all but first element (program)

	var sum float64 = 0 // var to hold sum of numbers
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err) // if err converting string, log and exit
		}
		sum += number
	}

	n := float64(len(arguments)) // length of the slice can be used
	fmt.Printf("Average: %0.2f\n", sum / n) // calculate the average and print it
}