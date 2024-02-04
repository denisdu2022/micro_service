package client_proxy

import (
	"net/rpc"
	"new_rpc/hanlder"
)

type HelloServiceStub struct {
	*rpc.Client
}

//在go语言中没有类对象,就没有初始化方法

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
