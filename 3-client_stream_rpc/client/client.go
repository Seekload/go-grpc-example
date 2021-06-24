package main

import (
	"context"
	pb "go-grpc-example/3-client_stream_rpc/proto"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

const Address string = ":8000"

func main() {
	// 1.连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	// 2.建立gRPC连接
	streamClient := pb.NewStreamServiceClient(conn)

	// 3.调用record，获取流
	stream, err := streamClient.Record(context.Background())
	if err != nil {
		log.Fatalf("call record err: %v", err)
	}

	for i := 0; i < 5; i++ {
		// 4.向流中发送数据
		err := stream.Send(&pb.StreamRequest{Data: strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	// 5.关闭流并获取返回的消息
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client stream close err: %v", err)
	}
	log.Printf("get from server,code:%v,value:%v", resp.GetCode(), resp.GetValue())
}
