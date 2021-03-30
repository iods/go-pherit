package main

import (
	"log"
	"net"

	"github.com/iods/go-pherit/internal/rpc/pb/alpha"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8150")
	if err != nil {
		panic(err)
	}

	s := alpha.Server{}

	grpcServer := grpc.NewServer()
	
	alpha.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 %v", err)
	}
}