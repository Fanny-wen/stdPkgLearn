package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/pb"
	"log"
)

func main() {
	// 1. 连接grpc服务
	conn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	// 2. 初始化 grpc客户端
	client := pb.NewSayNameClient(conn)

	var t = pb.Teacher{
		Age:  12,
		Name: "wuhudasima",
	}
	// 3. 调用远程服务
	teacher, err := client.SayHello(context.TODO(), &t)
	fmt.Printf("%+v--%v\n", *teacher, err)
}
