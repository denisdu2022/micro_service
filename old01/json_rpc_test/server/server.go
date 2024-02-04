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
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic("Server启动失败!")
	}

	//2.注册处理逻辑 handler
	err = rpc.RegisterName("HelloService", &HelloService{})

	for {
		//3.启动服务
		conn, _ := listener.Accept()

		//4.RPC 编码替换为json编码
		//go使用协程,异步的处理多个请求
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}
