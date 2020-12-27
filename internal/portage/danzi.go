package portage

import (
	"fmt"
	"reflect"
)

func DanziTypes() {
	danziCheckTypes()
	danziSwitchTypes()
}

// danziCheckTypes Returns the data type of the given inputs in Go.
func danziCheckTypes() {
	fmt.Println(reflect.TypeOf(12))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(1.0))
	fmt.Println(reflect.TypeOf("Green Danzi."))
}

func danziSwitchTypes() {
	alpha := 1
	fmt.Println(reflect.TypeOf(alpha))
	beta := 2
	fmt.Println(reflect.TypeOf(float64(beta)))
}