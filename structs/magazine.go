package magazine

type Person struct {
	Name        string
	Age         int
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}