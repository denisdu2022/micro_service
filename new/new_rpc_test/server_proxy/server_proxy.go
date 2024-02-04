package server_proxy

import (
	"net/rpc"
	"new_rpc_test/handler"
)

//需要解耦   --关心的是函数  鸭子类型

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
