/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/


package datatypes

import (
	"fmt"
	"strings"
)

// ChannelTypes Exportable function that contains callable examples
//    of using variables in Go.
func ChannelTypes() {
	title := "Channel: Variables in Go."
	fmt.Println(strings.Title(title))
	channelVars()
}

// channelVars Prints multiple variables of different types in Go.
func channelVars() {
	// declaring the variables
	//var name string
	//var length, width float64
	//var built int

	// assign on one line
	var keeper string = "John Langland"

	// omit the type and Go will use the type of the var
	var depth = 18

	// assigning the variables
	name := "Portage Point Pier"
	length, width := 150.0, 7.0
	built := 1891

	// zero values (set to false if not declared, 0 if int)
	var lighthouse bool

	// print the variables
	fmt.Println("Channel Name:", name)
	fmt.Println("Built In:", built)
	fmt.Println("Channel Length, Width (m):", length, width)
	fmt.Println("Channel Depth (ft):", depth)
	fmt.Println("Has Lighthouse:", lighthouse)
	fmt.Println("Lighthouse Keeper:", keeper)
}

