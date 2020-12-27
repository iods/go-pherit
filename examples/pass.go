package main
// Everything to the end of the line is a comment
// Additional lines require another `//` marker

/* Everything until closing marker is comment
   This includes new lines

   And blank lines */

/* Multiline comments can be single lined too */
import (
	"fmt"   // format
	"github.com/thedarksociety/go-pherit/pkg/keyboard"
	"log"
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

	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println(status)
}
