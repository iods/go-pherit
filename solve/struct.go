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

type car struct {
	name string
	topSpeed float64
}

type part struct {
	description string
	count int
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func doublePack(p *part) {
	p.count *= 2
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
	fmt.Println()
	// new exercise

	var mustang car
	mustang.name = "Ford Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)

	var fuses part
	fuses.description = "Fuses for the car."
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}