package main

import (
	"fmt"
)

func funkA(a int, b int) {
	fmt.Println(a + b)
}

func funkB(a int, b int) {
	fmt.Println(a * b)
}

func funkC(a bool) {
	fmt.Println(!a)
}

func funkD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Exercise for Functions in Go.")
	funkA(6, 6) // 12
	funkB(3, 4) // 12
	funkC(true) // false
	funkD("$", 4) // $$$$
	funkA(5, 6) // 11
	funkB(5, 6) // 30
	funkC(false) // true
	funkD("lol", 3) // lollollol
}
