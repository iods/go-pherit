package main

import "fmt"

type Gallons float64
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64 // satisfy the stringer interface
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}


func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	gas := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons.\n", gas, gas.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons.\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f milliliters.\n", milk, milk.ToMilliliters())
	fmt.Printf("%0.3f gallons equals %0.3f liters.\n", milk, milk.ToLiters())


	fmt.Println(gas)
	fmt.Println(water)
	fmt.Println(milk)
}
