/*
 Project:   go-pherit
 Copyright: (c) 2020-2021, Rye Miller
*/

package main

import "fmt"

/*

The Decorator Pattern

A decorator allows you to add functionality at runtime rather than compile time.

This means you do not have to manually write new code to extend and object's feature
set, you just use the decorator pattern to add functionality.

 1. Solves the issue of becoming restricted by an extensive sub-class architecture.
 2. Rather than inheritance, a decorator uses smaller classes (Go structs) to wrap functionality
 3. Objects can share structs rather than having to define individual classes

First, a component interface requires an action method to be implemented. It is done by
both concrete and decorator objects. A component may have zero to many decorators, but
a decorator cannot exist without a component.

*/


func main() {

	theBombPizza := &anchovyTopping{
		pizza: &artichokeTopping{
			pizza: &tomatoTopping{
				pizza: newCheesePizza(),
			},
		}}

	pepperoniPizza := &pepperoniTopping{
		pizza: newCheesePizza(),
	}

	fmt.Printf("Price of the bomb is %.2f with %d calories.\n", theBombPizza.getPrice(), theBombPizza.getCalories())
	fmt.Printf("Price of the pepperoni is %.2f with %d calories.\n", pepperoniPizza.getPrice(), pepperoniPizza.getCalories())
}