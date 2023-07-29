package main

import (
	"context"
	"grpc-test-server/chat"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

const iterations = 1000

func main() {
	test_gropc()
	// test_rest()
}

func test_rest() {
	for i := 0; i < iterations; i++ {
		_, err := http.Get("http://localhost:8080/SayHello")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func test_gropc() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Sent{}

	for i := 0; i < iterations; i++ {
		_, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling endpoint %v", err)
		}
	}
}
