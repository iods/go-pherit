package alpha

import (
	"context"
	"log"
)

type Server struct {} // Server for gRPC

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived message body from client %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}