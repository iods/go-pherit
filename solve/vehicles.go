/*
Output:
Speeding Up
Turning Left
Stopping
Turning Right
*/
package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding Up.")
}

func (c Car) Brake() {
	fmt.Println("Stopping.")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up.")
}

func (t Truck) Brake() {
	fmt.Println("Stopping.")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Mercedes G500")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Dodge Ram")
	vehicle.Brake()
	vehicle.Steer("right")
}