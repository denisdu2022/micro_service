package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	//1. 实例化一个server
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		return
	}
	//2. 注册处理逻辑handle
	_ = rpc.RegisterName("HelloService", &HelloService{})

	//3. 启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候

	//4. 绑定服务
	rpc.ServeConn(conn)

	//是否可以跨语言调用? 1.go语言的rpc的序列化和反序列化协议是Gob  2.能否替换成常见的序列化

}
