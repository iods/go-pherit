/*
Output:
Price is 100 dollars.
Tax is 8 dollars.
Total cost is 108 dollars.
120 dollars available.
Within budget? true
 */
package main

import (
	"fmt"
)

func main() {
	var price = float64(100)
	fmt.Println("Price is", price, "dollars.")

	var taxRate = 0.08
	var tax = price * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total = price + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds = float64(120)
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= availableFunds)
}
