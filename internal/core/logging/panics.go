package main

import "fmt"

func calmDown() {
	recover()
	fmt.Println("Nah, we cool.")
}

func freakOut() {
	defer calmDown()
	panic("oh, shit. sugar we are going down.")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally.")
}