package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 33)
	fmt.Printf("That will be $%f please.\n", 0.23 * 5)

	// letter following the % sign indicates which verb to use
	fmt.Printf("Print a float: %f\n", 3.14) // f for float
	fmt.Printf("Print an integer: %d\n", 300) // d for decimal int
	fmt.Printf("Print a string: %s\n", "string") // s for string
	fmt.Printf("Print a Boolean: %t\n", false) // t for true
	fmt.Printf("Print any value: %v %v %v\n", 1.2, "\t", true) // v for value based on whatever
	fmt.Printf("Print any formatted value: %#v %#v %#v\n", 1.2, "\t", true) // formatted value
	fmt.Printf("Print the type of a value: %T %T %T\n", 1.2, "\t", true) // T for Type
	fmt.Printf("Print a percent sign %%100\n") // still need percent signs

	fmt.Printf("%12s | %s \n", "Product", "Cost")
	fmt.Println("------------------------------")
	fmt.Printf("%12s | %4d\n", "Goldfish", 49)
	fmt.Printf("%12s | %4d\n", "Junior Mints", 79)
	fmt.Printf("%12s | %4d\n", "Milk", 250)

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456) // rounded three places
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456) // rounded two places
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456) // rounded one place
	fmt.Printf("%%.1f: %.1f\n", 12.3456) // rounded one place, no padding
	fmt.Printf("%%.2f: %.2f\n", 12.3456) // rounded two places, no padding
}
