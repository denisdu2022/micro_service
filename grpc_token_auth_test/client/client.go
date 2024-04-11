package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_token_auth_test/proto"
)

//PerRPCCredentials 接口的方法

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	//为了实现传输安全
	return false
}

func main() {

	//client拦截器函数

	//第一种方法
	//interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	//开始时间
	//	start := time.Now()
	//
	//	//使用metadata
	//	md := metadata.New(map[string]string{
	//		"appid":  "101010",
	//		"appkey": "i am key",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时: %s\r\n", time.Since(start))
	//	return err
	//}

	//第二种方法,使用WithPerRPCCredentials()
	//i := grpc.WithPerRPCCredentials(customCredential{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))

	conn, err := grpc.Dial(":8080", opts...) //opts... 不定长参数

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//获取客户端
	c := proto.NewGreeterClient(conn)

	//调用
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "libai"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
