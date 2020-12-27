/*
Golph Calculates the average of several numbers.
*/
package main

import (
	"fmt"
	"github.com/iods/go-datafile"
	"log"
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

func countCourses() {
	lines, err := datafile.GetStrings("courses.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Times Played at %s: %d\n", name, count)
	}
}

func main() {
	//arguments := os.Args[1:] // get slice of strings, all but first element (program)
	//
	//var numbers []float64 // var to hold sum of numbers
	//for _, argument := range arguments {
	//	number, err := strconv.ParseFloat(argument, 64)
	//	if err != nil {
	//		log.Fatal(err) // if err converting string, log and exit
	//	}
	//	numbers = append(numbers, number)
	//}
	//
	//fmt.Printf("Average: %0.2f\n", average(numbers...)) // calculate the average and print it

	//lines, err := datafile.GetStrings("courses.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(lines)
	//fmt.Println("-----------------------------")
	fmt.Println("Courses Played:")

	countCourses()

	// create an empty map
	empty := map[string]float64{}
	fmt.Println(empty)

	//var ranks map[string]int
	//ranks = make(map[string]int)
	//
	ranks := make(map[string]int)

	ranks["hard"] = 1
	ranks["medium"] = 2
	ranks["easy"] = 3

	courses := map[string]int{
		"Buffalo Ridge": ranks["hard"],
		"Vail Golf Course": ranks["medium"],
		"The Ridge At Castle Pines": ranks["hard"],
		"Riverdale Dunes": ranks["easy"],
		"Interlocken": ranks["hard"],
		"Indian Hills": ranks["easy"],
	}

	for name, rank := range courses {
		fmt.Println(name, rank)
	}
}