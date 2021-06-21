package main

import (
	pb "go-grpc-example/2-server_stream_rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

// 定义我们的服务
type StreamService struct{}

// 实现List方法
func (s *StreamService) List(req *pb.SimpleRequest, srv pb.StreamService_ListServer) error {
	for i := 0; i < 5; i++ {
		// 向流中发送消息，默认每次发送消息大小为 math.MaxInt32 byte
		err := srv.Send(&pb.StreamResponse{
			Code:  int32(i),
			Value: req.Data,
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	// 1.监听端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("listener err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 2.实例化gRPC服务端
	// 默认单次接受消息大小为 1024*1024*4 字节（4M），发送大小为 math.MaxInt32 字节
	grpcServer := grpc.NewServer()

	// 3.注册我们实现的服务 StreamService
	pb.RegisterStreamServiceServer(grpcServer, &StreamService{})

	// 4.启动gRPC服务端
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}
