package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([5]float64, error) {
	var numbers [5]float64

	file, err := os.Open(filename) // open the provided filename
	if err != nil {
		return numbers, err
	}

	i := 0 // tracks the array index to assign to
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) // convert line by line string to float
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close() // close for resources
	if err != nil {
		return numbers, err // if there was an error closing it, report it
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err() // and another one, and another one
	}

	return numbers, nil // finally return the array of numbers and nil error
}
