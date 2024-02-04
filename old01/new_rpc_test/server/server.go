package main

import (
	"net"
	"net/rpc"
	"new_rpc/hanlder"
	"new_rpc/server_proxy"
)

func main() {
	//1. 实例化一个server
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic("服务启动失败!")
	}

	//2. 注册处理逻辑handle
	_ = server_proxy.RegisterHelloService(&hanlder.NewHelloService{})

	for {
		//3.开启监听
		conn, _ := listener.Accept()

		//绑定服务
		go rpc.ServeConn(conn)
	}

}
