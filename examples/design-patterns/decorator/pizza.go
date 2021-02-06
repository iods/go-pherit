/*
 Project:   go-pherit
 Copyright: (c) 2020-2021, Rye Miller
*/

package main

// pizza Interface with getPrice() and getCalories() methods to serve as the component.
type pizza interface {
	getPrice()    float64
	getCalories() int
}

// cheesePizza Struct representing the concrete component of a decorator.
type cheesePizza struct {
	price    float64
	calories int
}

// newCheesePizza Returns a reference to a cheesePizza acting as a Constructor.
func newCheesePizza() *cheesePizza {

	return &cheesePizza{
		price: 10.00,
		calories: 8150,
	}
}

// getPrice Implements pizza interface by passing reference to a cheesePizza.
func (c *cheesePizza) getPrice() float64 {
	return c.price
}

func (c *cheesePizza) getCalories() int {
	return c.calories
}