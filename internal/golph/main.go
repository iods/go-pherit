package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/structs"
)

func main() {
	var c structs.Course
	c.Name = "Pebble Beach Golf Links"
	c.Location = "Pacific Surf"
	c.Founded = 1919
	c.Par = 72
	c.Points = 67.5261
	c.Rank = 7
	c.Rating = 4.7

	fmt.Println(c.Name)
	fmt.Println(c.Location)
	fmt.Println(c.Founded)
	fmt.Println(c.Par)
	fmt.Println(c.Points)
	fmt.Println(c.Rank)
	fmt.Println(c.Rating)

	var p structs.Player
	p.Name = "Rye Miller"
	p.Handicap = 15.4

	fmt.Println(p.Name, p.Handicap)

}