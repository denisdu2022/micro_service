package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_interpretor/proto"
	"time"
)

func main() {

	//拨号连接

	//client拦截器函数
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//开始时间
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时: %s\r\n", time.Since(start))
		return err
	}

	////client拦截器 :一元
	//第一种方法
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//
	//conn, err := grpc.Dial(":8080", grpc.WithInsecure(), opt)

	//第二种方法
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(":8080", opts...) //opts... 不定长参数

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//获取客户端
	c := proto.NewGreeterClient(conn)

	//使用metadata
	//第一种方法
	md := metadata.New(map[string]string{
		"name":     "libai",
		"password": "12345678",
	})

	//第二种方法
	//md := metadata.Pairs("timestamp", time.Now())

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//调用
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "libai"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
