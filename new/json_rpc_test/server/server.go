package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}

func main() {
	//1.实例化一个server
	listener, err := net.Listen("tcp", ":8000")
	//错误处理
	if err != nil {
		return
	}
	//2. 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	for {
		//3.启动服务
		conn, _ := listener.Accept()
		//绑定RPC服务,使用json编码
		//go 使用协程,异步处理,接收多个请求
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}
