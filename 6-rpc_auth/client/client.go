package main

import (
	"context"
	pb "go-grpc-example/5-security/proto"
	"google.golang.org/grpc"
	"log"
)

const Address = ":8000"

type Token struct {
	AppId     string
	AppSecret string
}

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_id": t.AppId, "app_secret": t.AppSecret}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return false
}

func main() {

	// Token token认证
	token := Token{
		AppId:     "grpc_auth",
		AppSecret: "123457",
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(&token),
	}

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		log.Fatalf("dial conn err: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewSimpleClient(conn)

	req := pb.SimpleRequest{
		Data: "seekload",
	}
	resp, err := grpcClient.GetSimpleInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("resp err: %v", err)
	}
	log.Printf("get from client,code: %v,value: %v", resp.Code, resp.Value)

}
