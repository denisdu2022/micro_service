package main

import (
	"fmt"
	"net"
	"net/rpc"
	"newhelloworld/hanlder"
	"newhelloworld/server_proxy"
)

func main() {

	//1.启动监听
	listener, err := net.Listen("tcp", ":8800")
	//错误处理
	if err != nil {
		fmt.Println("监听启动失败!")
		panic(err)
	}

	//3.注册处理逻辑 handler

	_ = server_proxy.RegisterHelloService(&hanlder.NewHelloService{})

	//循环接收
	for {
		//2.阻塞接收
		conn, err := listener.Accept()
		//错误处理
		if err != nil {
			fmt.Println("连接失败")
			panic(err)
		}

		//4.绑定服务 使用编解码 接收json-rpc实例化conn
		//请求会一个一个处理
		//rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//使用携程(异步处理),同时处理多个请求并发
		go rpc.ServeConn(conn)
	}

}
