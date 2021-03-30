package main

import (
	"context"
	"log"

	"github.com/iods/go-pherit/internal/rpc/pb/alpha"
	"google.golang.org/grpc"
)

func main() {

	var cc *grpc.ClientConn
	cc, err := grpc.Dial("localhost:8150", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connnect: %v", err)
	}
	defer cc.Close()

	c := alpha.NewChatServiceClient(cc)
	message := alpha.Message{
		Body: "Whassssssssssssup from the client.",
	}
	response, err := c.SayHello(context.Background(), &message)

	log.Printf("Response from Server: %s", response.Body)
}