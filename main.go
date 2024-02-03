package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	helloworldpb "grpctest/proto/helloworld"
	"log"
	"net"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	fmt.Println("got the request")
	return &helloworldpb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &server{})

	log.Println("gRPC server is running on :9090")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
