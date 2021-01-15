package gadget

import "fmt"

// first define an interface type
type Dwight interface {
	MethodWithoutParameters()      // a type satisfies this interface if it has this method
	MethodWithParameters(float64)
	MethodWithReturnValue() string // returns a string value
}

type Beats int

func (b Beats) MethodWithoutParameters() {
	fmt.Println("Method without params was called.")
}

func (b Beats) MethodWithParameters(f float64) {
	fmt.Printf("Method with param %0.2f was called.\n", f)
}

func (b Beats) MethodWithReturnValue() string {
	return "Method with return value was called."
}

func (b Beats) MethodNotInInterface() {
	fmt.Println("Method not in the interface was called.")
}