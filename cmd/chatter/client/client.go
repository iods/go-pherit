package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	client, _ := rpc.Dial("tcp", "localhost:8150")

	var reply = new(*string)

	err := client.Call("ChatterService.Hello", "Rye", reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(*reply)
}