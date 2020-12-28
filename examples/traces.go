package main

import "fmt"

func three() {
	panic("this call stack is tooooooo much for me.")
}

func two() {
	three()
}

func one() {
	defer fmt.Println("deferred in one.")
	two()
}

func main() {
	one()
}
