package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"grpc-helloworld/proto" // 替换为你的模块路径

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResp, error) {
	tSex := strconv.FormatInt(int64(req.Sex), 10) // 需要将 int 强制转换为 int64
	return &proto.HelloResp{Message: "Hello, " + req.Name + ", sex is " + tSex + ", happy!" + ", tags:" + strings.Join(req.Tags, ",") + ", home_num is " + req.Home.HomeNum}, nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	log.Println("Server is running on port :50051")
	// 启动服务
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
