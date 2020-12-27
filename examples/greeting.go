package main

import (
	"github.com/thedarksociety/go-pherit/pkg/greeting"
	"github.com/thedarksociety/go-pherit/pkg/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
}