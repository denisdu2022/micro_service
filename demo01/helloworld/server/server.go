package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//类

type HelloService struct {
}

//类的方法
//request传入参数 reply是传出参数

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "你好：" + request
	return nil
}

func main() {

	//1. 启动监听  实例化一个server
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("监听启动失败>>> ", err)
	}

	//2.注册处理逻辑 handel
	//注册RPC服务
	_ = rpc.RegisterName("HelloService", &HelloService{})

	//阻塞建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("conn err >>> ", err)
	}

	defer conn.Close()

	fmt.Println("net/rpc服务已启动...")
	//3.绑定服务
	rpc.ServeConn(conn)

}
