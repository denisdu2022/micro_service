package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"metadata_test/proto"
	"net"
	"reflect"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

//context.Context 是用来解决协程超时的

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	//取出metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}

	//for key, val := range md {
	//	fmt.Println(key, val)
	//}
	/*
		//输出很多内容
		get metadata error
		user-agent [grpc-go/1.63.2]
		name [libai]
		:authority [localhost:8080]
		content-type [application/grpc]
		password [12345678]
	*/

	//精确取值,取出name
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		//输出的是slice
		fmt.Println(reflect.TypeOf(nameSlice))
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	return &proto.HelloReply{
		Message: "Hello: " + request.Name,
	}, nil
}

func main() {
	//1.实例化grpc server
	g := grpc.NewServer()

	//2.注册 handler
	proto.RegisterGreeterServer(g, &Server{})

	//3.启动服务
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen:" + err.Error()) //err.Error() 具体的错误打印出来
	}

	//绑定服务
	err = g.Serve(listener)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
