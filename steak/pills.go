package main

import (
	"fmt"
	"strings"
	"time"
)

/*
auJus Returns an example of using integers in an array.
*/
func auJus() {
	primes := [5]int{2, 3, 5, 7, 11}
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[2])
}

/*
rare Returns an example of using time values within an array.
*/
func rare() {
	var dates [3]time.Time
	dates[0] = time.Unix(12578940090, 0)
	dates[1] = time.Unix(12448940090, 0)
	dates[2] = time.Unix(12578955530, 0)
	fmt.Println(dates[1])
}

/*
mediumRare Returns an example of zero values in arrays by setting one of integer counters.
*/
func mediumRare() {
	var counters [3]int
	counters[0]++
	fmt.Println(counters[0])
	counters[0]++
	counters[2]++
	fmt.Println(counters[0])
	fmt.Println(counters[1])
	fmt.Println(counters[2])
}

/*
main Arrays are like a pill box, each element has it's own index, or slot.
*/
func main() {
	title := "Steak: Arrays in Go."
	fmt.Println(strings.Title(title))

	var notes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[3])
	auJus()
	rare()
	mediumRare()
}