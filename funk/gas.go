package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64
type Hello string
type Number int

func (h Hello) Hello() {
	fmt.Printf("Hey, there, you %s. It worked.\n", h)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n) + otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n) - otherNumber)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(3.4)

	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))

	fmt.Printf("Car Fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus Fuel: %0.1f liters\n", busFuel)

	test := Hello("fucker")
	test.Hello()

	mil := Milliliters(422.0)
	fmt.Println(mil.ToGallons())

	mils := Milliliters(709.3)
	fmt.Println(mils.ToGallons())

	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
