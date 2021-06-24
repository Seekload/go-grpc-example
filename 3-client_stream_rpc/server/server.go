package main

import (
	pb "go-grpc-example/3-client_stream_rpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

// 定义我们的服务
type StreamService struct{}

// 实现 Record 方法
func (s *StreamService) Record(srv pb.StreamService_RecordServer) error {
	for {
		// 从流中获取消息
		req, err := srv.Recv()
		if err == io.EOF {
			// 发送数据并关闭
			return srv.SendAndClose(&pb.SimpleResponse{
				Code:  1,
				Value: "ok",
			})
		}
		if err != nil {
			return err
		}
		log.Printf("get from client:%v", req.Data)
	}
}

func main() {
	// 1.监听端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("listener err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	// 2.创建gRPC服务端实例
	grpcServer := grpc.NewServer()

	// 3.注册我们实现的服务 StreamService
	pb.RegisterStreamServiceServer(grpcServer, &StreamService{})

	// 4.启动gRPC服务端
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}
