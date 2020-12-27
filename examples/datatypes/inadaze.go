package datatypes

import (
	"fmt"
	"strings"
)

// InadazeTypes Returns an exported function for representing
//   various number and boolean types.
func InadazeTypes() {
	title := "Inadaze: Number types in Go."
	fmt.Println(strings.Title(title))
	inadazeOperators()
	inadazeAverage()
}

// inadazeOperators Returns various arithmetic operator usage examples in Go.
func inadazeOperators() {
	// Addition/sum operator
	fmt.Println("Sum (+):", 1 + 1)
	fmt.Println("Sum float64 (+):", 1 + 1.5)
	// Subtraction/difference operator
	fmt.Println("Difference (-):", 2 - 1)
	// Multiplication/product operator
	fmt.Println("Product (*):", 8 * 8)
	fmt.Println("Product (*):", 8 * -4.0) // -32.0 not -32
	// Division/quotient operator
	fmt.Println("Quotient (/):", 4 / 2)
	// Remainder operator (must be int)
	fmt.Println("Remainder (%):", 5 % 2)
	// Negation operator
	fmt.Println(- (-2))
	fmt.Println(- -2)
}

// inadazeAverage Returns the average age of the boats we have in Michigan.
func inadazeAverage() {
	var (
		slick   = 15
		inadaze = 19
		average float64
	)
	average = float64(slick+inadaze) / 2
	fmt.Println(average)
}
