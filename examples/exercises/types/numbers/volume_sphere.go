/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
Exercise: Find the volume of a sphere.
1. Get the radius from the command-line
2. Convert it to a float
3. Calculate the volume of a sphere

Output:
go run main.go 10
4188.79

go run main.go .5
0.52

Notes:

*/

func main() {
	var radius, volume float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	volume = ((4 * math.Pi) * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> area: %.2f\n", radius, volume)
}