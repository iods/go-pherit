/*
Exercise:

1. Print the following

Output:
I started with 10 apples.
Some jerk ate 4 apples.
There are 6 apples left.

Notes:

Bonus:
  * Print using Println and Printf

*/
package main

import "fmt"

func main() {
	var originalCount = 10
	var eatenCount = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount - eatenCount, "apples left.")

	fmt.Printf("I started with %d apples.\n", originalCount)              // bonus
	fmt.Printf("Some jerk ate %d apples.\n", eatenCount)                  // bonus
	fmt.Printf("There are %d apples left.\n", originalCount - eatenCount) // bonus
}