package main

import (
	"context"
	pb "go-grpc-example/5-security/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const Address string = ":8000"

func main() {
	// 为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "localhost")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	var DialOptions = []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	// 连接服务端
	conn, err := grpc.Dial(Address, DialOptions...)
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	// 创建gRPC客户端
	grpcClient := pb.NewSimpleClient(conn)

	res := pb.SimpleRequest{
		Data: "seekload",
	}
	resp, err := grpcClient.GetSimpleInfo(context.Background(), &res)
	if err != nil {
		log.Fatalf("stream get from server err: %v", err)
	}
	log.Printf("get from server,code: %v,value: %v", resp.Code, resp.Value)
}
