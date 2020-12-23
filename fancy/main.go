package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/structs"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Rye Miller"}
	subscriber.Street = "1469 Deer Blvd."
	subscriber.City = "Avon"
	subscriber.State = "CO"
	subscriber.PostalCode = "81620"

	employee := magazine.Employee{Name: "Tiffany Creamer"}
	employee.Street = "88 E Way Blvd."
	employee.City = "Denver"
	employee.State = "CO"
	employee.PostalCode = "80238"

	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Zip:", subscriber.PostalCode)

	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Zip:", employee.PostalCode)
}
