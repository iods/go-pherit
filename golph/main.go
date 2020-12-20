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

/*
twoInt Returns a slice using a variadic function.
*/
func twoInt(values ...int) { // expected two params
	fmt.Println(values)
}

/*
severalStrings Returns an empty slice when no arguments are received.
*/
func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func average(numbers ...float64) float64 {
	var sum float64 = 0 // var to hold the sum of the args
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:] // get slice of strings, all but first element (program)

	var numbers []float64 // var to hold sum of numbers
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err) // if err converting string, log and exit
		}
		numbers = append(numbers, number)
	}

	fmt.Printf("Average: %0.2f\n", average(numbers...)) // calculate the average and print it
}