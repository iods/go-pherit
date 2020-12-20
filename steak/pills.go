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
	// this is what happens when you don't reassign it back
	// to the original var
	prime := []int{
		2,
		3,
		5,
		7,
		11,
	}

	primes := append(prime, 13)
	primer := append(primes, 15)
	primed := append(primer, 17)

	fmt.Println(prime)
	fmt.Println(primes)
	fmt.Println(primer)
	fmt.Println(primed)

	primed[0] = 4

	fmt.Println(prime)
	fmt.Println(primes)
	fmt.Println(primer)
	fmt.Println(primed)
}

/*
rare Returns an example of using time values within an array.
*/
func rare() {
	var dates [3]time.Time
	dates[0] = time.Unix(12578940090, 0)
	dates[1] = time.Unix(12448940090, 0)
	dates[2] = time.Unix(12578955530, 0)

	for i := 0; i < 3; i ++ {
		fmt.Println(i, dates[i])
	}
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
medium Returns an example of checking the array length in a for loop (resist the panic!)
*/
func medium() {
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"} // no need to use make() slice literals are made automatically
	fmt.Println(len(letters))

	for i := 0; i < len(letters); i++ {
		fmt.Println(i, letters[i])
	}
	// or just letter
	for _, letter := range letters {
		fmt.Println(letter)
	}

	sliceOne := letters[0:3]
	sliceTwo := letters[2:5]

	letters[2] = "X" // changing the underlying array will change the slice

	fmt.Println(sliceOne, sliceTwo)
}

/*
mediumWell Returns an example of using for...range loops with arrays. Stupid simple amazing.
*/
func mediumWell() {
	var words = [6]string{"alpha", "beta", "charlie", "delta", "echo", "foxtrot"}
	others := words[0:3]
	for index, word := range words {
		fmt.Println(index, word)
	}

	for i, other := range others {
		fmt.Println(i, other)
	}
}

/*
well Returns an example of using a blank identifier in a for...range loop.
*/
func well() {
	var words = [6]string{"alpha", "beta", "charlie", "delta", "echo", "foxtrot"}
	i, j := 1, 5
	others := words[i:j]
	for index, _ := range words {
		fmt.Println(index)
	}
	for i, other := range others {
		fmt.Println(i, other)
	}
}

/*
wellDone Returns an example of a function using a slice to set an array variable.
*/
func wellDone() {
	var meat []string // same as the syntax for declaring array vars, just no size
	var rating []int

	meat = make([]string, 7) // does not create on own, use make()
	meat[0] = "filet"
	meat[1] = "flank"
	meat[2] = "philly"

	rating = make([]int, 7)
	rating[0] = 1
	rating[1] = 2
	rating[2] = 3

	fmt.Println("Meat:", meat[1], " Rating:", rating[2])
}


/*
main Arrays are like a pill box, each element has it's own index, or slot.
*/
func main() {
	title := "Steak: Arrays in Go."
	fmt.Println(strings.Title(title))

	var notes = []string{"do", "re", "mi", "fa", "so", "la", "ti", "da"}

	index := 3
	fmt.Println(notes[0], notes[1], notes[index])
	auJus()
	rare()
	mediumRare()
	medium()
	mediumWell()
	well()
	wellDone()
}