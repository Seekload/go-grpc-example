package main

import (
	"context"
	pb "go-grpc-example/1-simple_rpc/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	Address string = ":8000"
)

func main() {
	// 1.创建于gRPC服务端的连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial conn err: %v", err)
	}
	defer conn.Close()

	// 2.创建grpc客户端
	client := pb.NewSimpleClient(conn)

	// 3.调用服务端提供的服务
	req := pb.SimpleRequest{
		Data: "Hello,Server",
	}
	resp, err := client.GetSimpleInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("resp err: %v", err)
	}
	log.Printf("get from server,code: %v,value: %v", resp.GetCode(), resp.GetValue())

}
