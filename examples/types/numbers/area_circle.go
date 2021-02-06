/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Area of a Circle
1. Calculate the area of a circle from the given radius.

Output:
radius: 10 -> area: 314.16

Notes:
  * use math.Pi
  * area = πr² is the formula

*/

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %0.f -> area: %0.2f\n", radius, area)
}