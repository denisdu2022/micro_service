package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"metadata_test/proto"
)

func main() {

	//拨号连接
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//获取客户端
	c := proto.NewGreeterClient(conn)

	//使用metadata
	//第一种方法
	md := metadata.New(map[string]string{
		"name":     "libai",
		"password": "12345678",
	})

	//第二种方法
	//md := metadata.Pairs("timestamp", time.Now())

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//调用
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "libai"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
