package chat

import (
	"context"
)

type Server struct {
	Data []*Response
}

func (s *Server) SayHello(ctx context.Context, sent *Sent) (*Message, error) {
	// log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: s.Data}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
