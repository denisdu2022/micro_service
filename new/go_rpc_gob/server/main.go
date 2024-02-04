package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello: " + request
	return nil
}

func main() {
	//1.实例化一个server
	listener, err := net.Listen("tcp", ":8000")
	//错误处理
	if err != nil {
		return
	}
	//2.注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候
	//绑定rpc服务
	rpc.ServeConn(conn)

}
