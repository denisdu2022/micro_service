package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_validate_test/proto"
)

func main() {
	//grpc client 参数
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	//建立连接
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//grpc的client

	c := proto.NewGreeterClient(conn)
	//调用
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:    1000,
		Email: "test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("rspID: %d ", rsp.Id)
}
