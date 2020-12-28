package main

import (
	"fmt"
	"log"
)

/*
Socialize Prints a sequence of strings using the defer keyword.
*/
func Socialize() error {
	defer fmt.Println("Goodbye.")
	fmt.Println("Hello, There.")

	return fmt.Errorf("i don't want to talk")
	fmt.Println("Nice, weather, eh?")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}