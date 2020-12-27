package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	temp, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (temp - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius.\n", celsius)
}