package chatter

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// Run Establishes a simple server within Go by listening
// on a port for any incoming TCP connections. Then by importing
// the Google gRPC module, register the endpoints, and serve.
func Run() {

	connections, err := net.Listen("tcp", ":8150")
	if err != nil {
		log.Fatalf("failed to listen on port: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(connections); err != nil {
		log.Fatalf("failed to serve: %s\n", err)
	}
}

