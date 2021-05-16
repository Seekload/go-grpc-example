package main

import (
	"context"
	pb "go-grpc-example/1-simple_rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

// 定义我们的服务
type SimpleService struct{}

// 实现 GetSimpleInfo 方法
func (s *SimpleService) GetSimpleInfo(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := req.Data
	log.Println("get from client: ", data)
	resp := &pb.SimpleResponse{
		Code:  8888,
		Value: "grpc",
	}
	return resp, nil
}

func main() {

	// 1.监听端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}
	log.Println(Address, " net listening...")
	// 2.实例化gRPC服务端
	grpcServer := grpc.NewServer()

	// 3.注册我们实现的服务 SimpleService
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	// 4.启动gRPC服务端
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v",err)
	}

}
