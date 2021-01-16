package main

import (
	"flag"
	"fmt"
	"github.com/iods/go-pherit/pkg/eddie"
)

func main() {
	c := eddie.App{}
	c.Setup()

	flag.Parse()
	fmt.Println(c.GetMessage())
}
