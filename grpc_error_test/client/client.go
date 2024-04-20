package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"grpc_error_test/proto"
	"time"
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

	//超时时间  context.WithTimeout()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	//3.远程调用
	_, err = c.SayHello(ctx, &proto.HelloRequest{
		Name: "李白"})
	//错误处理
	if err != nil {
		//panic(err)
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	//fmt.Println("message: ", r.Message)

}
