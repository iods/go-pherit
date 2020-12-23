package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/structs"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Is Active:", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Rye Miller")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	sub2 := defaultSubscriber("Tiffany Creamer")
	printInfo(sub2)
}