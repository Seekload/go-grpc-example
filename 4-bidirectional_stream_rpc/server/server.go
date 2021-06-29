package main

import (
	pb "go-grpc-example/4-bidirectional_stream_rpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

// 定义我们的服务
type StreamService struct{}

// 实现 Record() 方法
func (s *StreamService) Record(srv pb.StreamService_RecordServer) error {
	n := 1
	for {
		// 接收数据
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("stream get from client err: %v", err)
			return err
		}
		// 发送数据
		err = srv.Send(&pb.StreamResponse{
			Code:  int32(n),
			Value: "This is the " + strconv.Itoa(n) + " message",
		})
		if err != nil {
			log.Fatalf("stream send to client err: %v", err)
			return err
		}
		n++
		log.Println("stream get from client: ", req.Data)
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

	// 2.实例化gRPC实例
	grpcServer := grpc.NewServer()

	// 3.注册我们的服务
	pb.RegisterStreamServiceServer(grpcServer, &StreamService{})

	// 4.启动gRPC服务端
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}
