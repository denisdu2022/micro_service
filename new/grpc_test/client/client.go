package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/proto"
)

func main() {

	//1.拨号连接
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	//错误处理
	if err != nil {
		panic(err)
	}
	//延迟注册关闭
	defer conn.Close()

	//2.获取client
	c := proto.NewGreeterClient(conn)
	//3.远程调用
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "李白"})
	//错误处理
	if err != nil {
		panic(err)
	}
	fmt.Println("message: ", r.Message)

}
