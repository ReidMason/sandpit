package main

import (
	"encoding/json"
	"grpc-test-server/chat"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/photos")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data []*chat.Response
	json.Unmarshal(body, &data)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{
		Data: data,
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
