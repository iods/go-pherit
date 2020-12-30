package main

import (
	"fmt"
)

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)

	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	myInt = 42
	fmt.Println(*myIntPointer)
}