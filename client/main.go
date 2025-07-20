package main

import (
	"context"
	"log"
	"time"

	"grpc-helloworld/proto" // 替换为你的模块路径

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到 gRPC 服务端
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := proto.NewGreeterClient(conn)

	// 设置请求上下文，超时 1 秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发送请求
	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "World Jam", Sex: 1, Tags: []string{"tag1", "tag2"}, Home: &proto.Home{HomeNum: "no.31"}})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	// 打印响应
	log.Printf("Response: %s", resp.Message)
}
