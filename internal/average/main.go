package main

import (
	"fmt"
	"github.com/iods/go-datafile"
	"github.com/iods/go-util"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err) // if there was an error opening that file, report and exit
	}

	fmt.Printf("Average: %0.2f\n", util.Average(numbers...))
}
