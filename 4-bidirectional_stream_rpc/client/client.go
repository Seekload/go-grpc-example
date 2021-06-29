package main

import (
	"context"
	pb "go-grpc-example/4-bidirectional_stream_rpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

const Address = ":8000"

func main() {
	// 1.连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	// 2.创建gRPC客户端
	grpcClient := pb.NewStreamServiceClient(conn)

	// 3.调用 Record() 方法获取流
	stream, err := grpcClient.Record(context.Background())
	if err != nil {
		log.Fatalf("call record err: %v", err)
	}

	for i := 0; i < 5; i++ {
		// 4.发送数据
		err := stream.Send(&pb.StreamRequest{
			Data: strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream send to server err: %v", err)
		}
		// 5.接收服务端发送过来的数据
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream get from server err: %v", err)
		}
		log.Printf("stream get from server,code:%v,value:%v", resp.GetCode(), resp.Value)
		time.Sleep(1 * time.Second)
	}
	// 6.关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("close stream err:%v", err)
	}
}
