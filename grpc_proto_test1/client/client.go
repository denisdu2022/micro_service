package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_proto_test1/proto"
	"time"
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
	//调用
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "libai",
		Url:  "baidu.com",
		//enum类型
		G: proto.Gender_MALE,
		//map类型
		Mp: map[string]string{
			"name":    "jack",
			"company": "sky",
		},
		//时间戳类型
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
