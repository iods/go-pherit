package main

import (
	"fmt"
	"strings"
)

type course struct {
	name string
	rating int
}

type golfer struct {
	name string
	handicap float64
}

func showCourse(c course) {
	fmt.Println("Course:", c.name)
	fmt.Println("Rating:", c.rating)
}

func showGolfer(g golfer) {
	fmt.Println("Name:", g.name)
	fmt.Println("Handicap:", g.handicap)
}

func main() {
	title := "Fancy: Structs in Go."
	fmt.Println(strings.Title(title))

	var g golfer
	g.name = "Rye Miller"
	g.handicap = 18.4
	fmt.Println(g.name)
	fmt.Println(g.handicap)

	var c course
	c.name = "Buffalo Ridge"
	c.rating = 2
	fmt.Println(c.name)
	fmt.Println(c.rating)

	showCourse(c)
	showGolfer(g)

	fmt.Printf("%s, with a %0.1f handicap, played at %s, a course ranked at %d.\n", g.name, g.handicap, c.name, c.rating)
}