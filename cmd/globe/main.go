package main

import (
	"github.com/thedarksociety/go-pherit/pkg/grape"
	"time"
)

func main() {

	a, b, c := "https://ryemiller.io", "https://golang.org", "https://golang.org/doc"

	grape.ResponseBody()

	go grape.ResponseSize(a)
	go grape.ResponseSize(b)
	go grape.ResponseSize(c)

	time.Sleep(5 * time.Second)
}
