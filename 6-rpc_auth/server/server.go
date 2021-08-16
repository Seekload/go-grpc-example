package main

import (
	"context"
	pb "go-grpc-example/6-rpc_auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

type SimpleService struct{}

func (s *SimpleService) GetSimpleInfo(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	// 检测Token是否有效
	if err := check(ctx); err != nil {
		return nil, err
	}
	data := req.Data
	log.Printf("get from client: %v", data)

	resp := pb.SimpleResponse{
		Code:  1,
		Value: "gRPC",
	}
	return &resp, nil
}

func main() {

	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}
	log.Println(Address, " net listening...")

	grpcServer := grpc.NewServer()

	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}

func check(ctx context.Context) error {
	// 从上下文章获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取token失败")
	}

	var (
		appId     string
		appSecret string
	)
	if value, ok := md["app_id"]; ok {
		appId = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appId != "grpc_auth" || appSecret != "123456" {
		return status.Errorf(codes.Unauthenticated, "Token无效，appId:%v,appSecret:%v", appId, appSecret)
	}
	return nil
}
