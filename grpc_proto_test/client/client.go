package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_proto_test/proto-back"
)

func main() {

	//拨号连接
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//获取客户端
	c := proto_back.NewGreeterClient(conn)
	//调用
	r, err := c.SayHello(context.Background(), &proto_back.HelloRequest{
		//URL和Name 编号反了
		//protobuf的编码
		/*
			json的格式:
			{
				"name":"bobby"
			}

			protobuf的格式:
			15libai29baidu.com
			客户端是反的,服务端解码时把顺序就会放反
		*/
		Name: "libai",
		Url:  "baidu.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
