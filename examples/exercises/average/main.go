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
		log.Fatal(err)
	}

	fmt.Printf("Average: %0.2f\n", util.Average(numbers...))
}
