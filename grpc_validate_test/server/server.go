package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_validate_test/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id: 1,
	}, nil
}

//grpc validate

type Validator interface {
	Validator() error
}

func main() {
	//拦截器
	var interceptor grpc.UnaryServerInterceptor
	//fmt.Println(reflect.TypeOf(interceptor))
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//继续处理请求
		if r, ok := req.(Validator); ok {
			if err := r.Validator(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}

	//grpc 的服务端参数
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	//启动监听服务
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	//grpc服务
	g := grpc.NewServer(opts...)
	//绑定grpc服务
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}

}
