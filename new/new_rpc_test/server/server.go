package main

import (
	"net"
	"net/rpc"
	"new_rpc_test/handler"
	"new_rpc_test/server_proxy"
)

func main() {

	//1.实例化一个server
	listener, err := net.Listen("tcp", ":8000")
	//错误处理
	if err != nil {
		return
	}
	//2.注册处理逻辑 handler
	err = server_proxy.RegisterHelloService(&handler.NewHelloService{})

	for {
		//3.启动服务
		conn, _ := listener.Accept()
		//绑定服务
		go rpc.ServeConn(conn)

	}

}
