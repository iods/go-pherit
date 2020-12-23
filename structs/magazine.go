package magazine

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Employee struct {
	Name       string
	Salary     float64
	Street     string
	City       string
	State      string
	PostalCode string
}

type Subscriber struct {
	Name       string
	Rate       float64
	Active     bool
	Street     string
	City       string
	State      string
	PostalCode string
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}