package main

import (
	"net"
	"net/rpc"
)

type ChatterService struct {}

func (s *ChatterService) Hello(request string, reply *string) error {

	*reply = "hello, " + request
	return nil
}

func main() {

	listener, _ := net.Listen("tcp", ":8150")
	_ = rpc.RegisterName("ChatterService", &ChatterService{})

	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}