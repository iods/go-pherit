/*
 Project:   go-pherit
 Copyright: (c) 2020-2021, Rye Miller
*/

package main

const (
	anchovyPrice = 2.50
	artichokePrice = 2.00
	pepperoniPrice = 1.75
	tomatoPrice = 0.75

	anchovyCalories = 350
	artichokeCalories = 150
	pepperoniCalories = 200
	tomatoCalories = 50
)

type anchovyTopping struct {
	pizza pizza
}

type artichokeTopping struct {
	pizza pizza
}

type pepperoniTopping struct {
	pizza pizza
}

type tomatoTopping struct {
	pizza pizza
}

func (a *anchovyTopping) getPrice() float64 {
	return a.pizza.getPrice() + anchovyPrice
}

func (a *anchovyTopping) getCalories() int {
	return a.pizza.getCalories() + anchovyCalories
}

func (a *artichokeTopping) getPrice() float64 {
	return a.pizza.getPrice() + artichokePrice
}

func (a *artichokeTopping) getCalories() int {
	return a.pizza.getCalories() + artichokeCalories
}

func (p *pepperoniTopping) getPrice() float64 {
	return p.pizza.getPrice() + pepperoniPrice
}

func (p *pepperoniTopping) getCalories() int {
	return p.pizza.getCalories() + pepperoniCalories
}

func (t *tomatoTopping) getPrice() float64 {
	return t.pizza.getPrice() + tomatoPrice
}

func (t *tomatoTopping) getCalories() int {
	return t.pizza.getCalories() + tomatoCalories
}