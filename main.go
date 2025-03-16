package main

import (
	"go-grpc/handler"
	"go-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot start server %s", err)
	}

	server := grpc.NewServer()
	service := &handler.BlogServer{}
	proto.RegisterBlogServiceServer(server, service)
	err = server.Serve(list)
	if err != nil {
		log.Fatalf("cannot serve requests: %s", err)
	}
}
