package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_test/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

//context.Context 是用来解决协程超时的

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
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
