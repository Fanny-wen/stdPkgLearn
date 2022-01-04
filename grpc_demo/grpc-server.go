package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_demo/pb"
	"log"
	"net"
)

// Children 定义类
type Children struct {
}

func (c Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "\tsay hello to you"
	t.Age += 1
	return t, nil
}

func main() {
	// 1. 初始化一个grpc对象
	grpcServer := grpc.NewServer()
	// 2. 注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	// 3. 设置监听, 指定 ip port
	lis, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	// 4. 启动服务
	_ = grpcServer.Serve(lis)
}
