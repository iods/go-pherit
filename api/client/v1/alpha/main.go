package main

import (
	"log"

	"github.com/iods/go-pherit/internal/rpc/services/alpha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8150", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %v", err)
	}
	defer conn.Close()

	c := alpha.NewServiceAlphaClient(conn)
	response, err := c.GetText(context.Background(), &alpha.Message{Body: "Hello, World, from the client."})
	if err != nil {
		log.Fatalf("error when calling %v", err)
	}

	log.Printf("Response from server: %s\n", response.Body)
}