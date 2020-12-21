/*
Exercise:
Produce the output:
Earrings: 79.99
Necklace: 89.99
Shirt: 39.99
Pants: 59.99
*/
package main

import "fmt"

func main() {
	// map using make
	jewelery := make(map[string]float64)

	jewelery["earrings"] = 79.99
	jewelery["necklace"] = 89.99

	fmt.Println("Earrings:", jewelery["earrings"])
	fmt.Println("Necklace:", jewelery["necklace"])

	// map using a literal
	clothing := map[string]float64{"shirt": 39.99, "pants": 59.99}
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}
