/*
What is Type Assertion?

As the name suggests they are used to assert that a variable is of some type.

Type assertions are used to check if the value held by the interface type variable
either implements the desired interface (satisfies) or is of a concrete type.

They can only take place on interfaces.

Type assertions are performed at runtime.

x.(T)
T is a concrete type or interface
One (asserted value) or two (ok) values can be returned

Why would I use one?
What exactly do they return?

x.(T) assert that x is not nil and that the value stored in x is type of T

Why would I use them?
 * to check that x is nil
 * to check if it is convertible (assert) to another type
 * convert (assert) it to another type

What exactly do they return?
 t := x.(T) => is of type T; if x is nil, it panics
 t, ok := x.(t) => if x is nil or not of type T => ok is false otherwise ok is true and t is of type T

vs. Type Conversion
*/
package main

import "fmt"

func check() {
	var x interface{} = "Foo"

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is", v)
	case bool, string:
		fmt.Println("x is a bool or a string.")
	default:
		fmt.Println("type unknowns.")
	}
}

func assert(i interface{}) {
	s := i.(int) // get the int value from i
	fmt.Println(s)
}

func main() {
	var s interface{} = 56
	assert(s)
}
