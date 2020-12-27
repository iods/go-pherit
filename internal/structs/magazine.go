package structs

type Person struct {
	Name        string
	Age         int
	Address
}

type Employee struct {
	Name        string
	Salary      float64
	Address
}

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}