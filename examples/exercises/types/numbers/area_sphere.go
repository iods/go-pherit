/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Area of a Sphere
1. Get the radius from the command line
2. Convert it to a float64
3. Calculate the surface area of a sphere

Output:
go run main.go 10
1256.64

go run main.go 54.2
36915.47

Notes:
  * use math.Pow

*/

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var radius, area float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	area = 4 * math.Pi * math.Pow(radius, 2)
	
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}