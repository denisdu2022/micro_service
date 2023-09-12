package server_proxy

import (
	"net/rpc"
	"newhelloworld/hanlder"
)

//关心的是函数,对RegisterHelloService 解耦合 鸭子模型

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
