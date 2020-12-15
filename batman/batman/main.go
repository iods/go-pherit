package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("I am Batman.")
	fmt.Println(runtime.NumCPU() + 1)
}