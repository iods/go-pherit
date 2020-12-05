package main

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/uno/packages/portage"
)

func main() {
	portage.Slick()
	fmt.Println(portage.Version())
	fmt.Println(portage.Cpus() + 4)
}