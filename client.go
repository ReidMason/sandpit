package main

import (
	"context"
	"grpc-test-server/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	reponse, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling endpoint %v", err)
	}

	log.Printf("Response from server: %s", reponse.Body)
}
