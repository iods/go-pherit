package main


import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/iods/go-pherit/internal/rpc/services/alpha"
	"google.golang.org/grpc"
)

type Server struct {}

func (s *Server) mustEmbedUnimplementedServiceAlphaServer() {
	panic("implement me")
}

func (s *Server) GetText(ctx context.Context, in *alpha.Message) (*alpha.Message, error) {
	log.Printf("Received message body from client: %s\n", in.Body)

	return &alpha.Message{Body: "Whasssssssssssssssssupppp from the server."}, nil
}


func main() {
	fmt.Println("The gRPC starter Pack")

	lis, err := net.Listen("tcp", fmt.Sprintf("%d", 8150))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	alpha.RegisterServiceAlphaServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}