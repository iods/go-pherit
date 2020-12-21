/*
Exercise:
Output:
Name: Max
Age: 5
*/
package main

import "fmt"

func main() {
	var pet struct{
		name string
		age int
	}
	pet.name = "Eddie"
	pet.age = 3

	fmt.Println("Name", pet.name)
	fmt.Println("Age:", pet.age)
}