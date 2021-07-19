package main

import (
	"context"
	pb "go-grpc-example/5-security/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

type SimpleService struct{}

func (s *SimpleService) GetSimpleInfo(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := req.Data
	log.Println("get from client: ", data)
	resp := pb.SimpleResponse{
		Code:  1,
		Value: "gRPC",
	}
	return &resp, nil
}

func main() {
	// 1.监听端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("listener err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	// 为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	// 2.实例化gRPC实例
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 3.注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	// 4.启动gRPC服务端
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}
