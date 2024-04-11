package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc_token_auth_test/proto"
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

	//interceptor函数
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")

		//取出metadata的值
		md, ok := metadata.FromIncomingContext(ctx)
		//fmt.Println(md)
		if !ok {
			//这里已经接触到grpc的错误处理了
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		//fmt.Println("------------------------------------------")
		var (
			appid  string
			appkey string
		)

		//判断appid 和 appkey
		if va1, ok := md["appid"]; ok {
			//返回的是[]string ,所以我们这里只取值第一个
			appid = va1[0]
		}

		if va1, ok := md["appkey"]; ok {
			//返回的是[]string ,所以我们这里只取值第一个
			appkey = va1[0]
		}

		fmt.Println(appid, appkey)

		if appid != "101010" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		res, err := handler(ctx, req)
		if err != nil {
			panic(err)
		}
		fmt.Println("请求已经完成")
		return res, err
	}

	//grpc一元方法拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	//1.实例化grpc server
	g := grpc.NewServer(opt)

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
