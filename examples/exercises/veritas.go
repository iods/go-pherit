package main

import (
	"fmt"
)

func Conditions() {

	if true {
		fmt.Println("true") // prints true
	}

	if false {
		fmt.Println("false") // empty (!false)
	}

	if !false {
		fmt.Println("!false") // !false
	}

	if true {
		fmt.Println("if true") // if true
	} else {
		fmt.Println("else")
	}

	if false {
		fmt.Println("if false")
	} else if true {
		fmt.Println("else if true") // else if true
	}

	if 12 == 12 {
		fmt.Println("12 == 12") // 12 == 12
	}

	if 12 != 12 {
		fmt.Println("12 != 12") // false or empty
	}

	if 12 > 12 {
		fmt.Println("12 > 12") // false or empty
	}

	if 12 >= 12 {
		fmt.Println("12 >= 12") // 12 >= 12
	}

	if 12 == 12 && 5.9 == 5.9 {
		fmt.Println("12 == 12 && 5.9 == 5.9") // 12 == 12 && 5.9 == 5.9
	}

	if 12 == 12 && 5.9 == 6.4 {
		fmt.Println("12 == 12 && 5.9 == 6.4") // false or empty
	}

	if 12 == 12 || 5.9 == 6.4 {
		fmt.Println("12 == 12 || 5.9 == 6.4") // 12 == 12 || 5.9 == 6.4
	}
}
