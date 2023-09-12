package client_proxy

import (
	"net/rpc"
	"newhelloworld/hanlder"
)

type HelloServiceStub struct { // 创建HelloServiceStud这个类
	*rpc.Client //要调用client的call方法
}

//在go语言中没有类 对象 ,就意味着没有初始化方法(初始化方法名一般都是New开头)

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{
		conn,
	}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
