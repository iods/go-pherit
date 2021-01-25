package main

import (
	"flag"
	"fmt"
	"github.com/iods/go-pherit/pkg/eddie"
)

func main() {

	eddie.Init()

	c := &eddie.Car{}

	for i := 0; i < 30; i++ {
		c.Model = append(c.Model, "Item " + fmt.Sprint(i))
	}

	fmt.Println(c.Model[5])
	fmt.Println(eddie.Flagname)
	fmt.Println(eddie.Strflag)
	fmt.Println(eddie.Ip)
	fmt.Println(flag.Args())
}
