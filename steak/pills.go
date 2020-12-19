package main

import (
	"fmt"
	"strings"
)

/*
main Arrays are like a pill box, each element has it's own index, or slot.
*/
func main() {
	title := "Steak: Arrays in Go."
	fmt.Println(strings.Title(title))

	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
}
