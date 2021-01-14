package User

import "github.com/iods/go-golph/examples/struct/Person"

type User struct {
	Person Person.Person
	Email  string
}