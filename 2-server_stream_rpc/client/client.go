package main

import (
	"context"
	pb "go-grpc-example/2-server_stream_rpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

const Address string = ":8000"

func main() {
	// 1.连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	// 2.创建grpc客户端
	grpcClient := pb.NewStreamServiceClient(conn)

	// 3.调用服务端提供的服务
	req := &pb.SimpleRequest{
		Data: "Hello,Server",
	}
	stream, err := grpcClient.List(context.Background(), req)
	if err != nil {
		log.Fatalf("call server err,%v", err)
	}
	for {
		// 4.处理服务端发送过来的流信息
		resp, err := stream.Recv()
		if err == io.EOF { // 流是否结束
			break
		}
		if err != nil {
			log.Fatalf("client get stream err:%v", err)
		}
		log.Printf("get from stream server,code:%v,value:%v", resp.GetCode(), resp.GetValue())
	}
}
