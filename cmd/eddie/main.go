package main

import (
	"fmt"
	"github.com/iods/go-pherit/pkg/eddie"
)

func main() {

	eddie.Init()

	c := &eddie.Car{}

	for i := 0; i < 15; i++ {
		c.Model = append(c.Model, "Item " + fmt.Sprint(i))
	}

	fmt.Println(c.Model[5])
	fmt.Println(eddie.Flagname)
}
