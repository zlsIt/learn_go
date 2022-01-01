package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "learn_go/grpc/proto/hello"
	"net"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatal("Failed to listen：%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, HelloService)
	grpclog.Info("listen on ", Address)
	s.Serve(listen)
}
