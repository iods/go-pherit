/*
Exercise:
Output:
Name: Max
Age: 5
---------------
Exercise: Code Magnets
Output:
Name: Alonzo Cole
Grade: 92.3
*/
package main

import "fmt"

type student struct {
	name string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var pet struct{
		name string
		age int
	}
	pet.name = "Eddie"
	pet.age = 3

	fmt.Println("Name", pet.name)
	fmt.Println("Age:", pet.age)
	fmt.Println()

	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3

	printInfo(s)
}