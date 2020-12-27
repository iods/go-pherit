package main

import "fmt"

type Whistle string

type Horn string

type Robot string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep.")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs.")
}

type NoiseMaker interface { // will represent any type w/ MakeSound method()
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Name of Whistle")
	toy.MakeSound()

	toy = Horn("Name of a Horn.")
	toy.MakeSound()

	play(Whistle("Another name bro."))
	play(Horn("A horn bro."))

	play(Robot("And another."))
}