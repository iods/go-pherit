package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use")
	flag.Parse()

	fmt.Printf("port = %d\n", port)
}
