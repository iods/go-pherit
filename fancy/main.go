package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/structs"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Rye Miller"}
	subscriber.HomeAddress.Street = "1469 Deer Blvd."
	subscriber.HomeAddress.City = "Avon"
	subscriber.HomeAddress.State = "CO"
	subscriber.HomeAddress.PostalCode = "81620"

	employee := magazine.Employee{Name: "Tiffany Creamer"}
	employee.HomeAddress.Street = "88 E Way Blvd."
	employee.HomeAddress.City = "Denver"
	employee.HomeAddress.State = "CO"
	employee.HomeAddress.PostalCode = "80238"

	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("Zip:", subscriber.HomeAddress.PostalCode)

	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("Zip:", employee.HomeAddress.PostalCode)
}
