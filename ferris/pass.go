package ferris

import (
	"bufio" // buffered reader
	"fmt"   // format
	"log"
	"os"
	"strconv"
	"strings" // strings package
)

/*
1. Allow the user to enter a number, store the number in a var
2. If the entered percentage is higher than 60% we set it to passing
*/
func PassFail() {
	title := "Ferris: Conditionals and Loops in Go."
	fmt.Println(strings.Title(title))

	// fmt.Print doesnt skip to a new line, we can grab input
	fmt.Print("Enter a grade: ")

	// read (receive and store) input from stdin
	// all input from the keyboard is read as a string.
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	var status string

	// scope is important, must be declared
	// cannot be :=
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println(status)
}
