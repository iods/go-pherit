package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err) // if there was an error opening that file, report and exit
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}
