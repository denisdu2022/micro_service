package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_error_test/proto"
	"net"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

//context.Context 是用来解决协程超时的

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	//grpc 中的错误处理
	//return nil, status.Error(codes.NotFound, "记录未找到")
	//return nil, status.Errorf(codes.NotFound, "记录未找到:%s", request.Name)

	//构建服务端超时
	time.Sleep(time.Second * 5)

	return &proto.HelloReply{
		Message: "helli",
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
