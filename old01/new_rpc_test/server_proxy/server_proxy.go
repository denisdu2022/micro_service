package server_proxy

import (
	"net/rpc"
	"new_rpc/hanlder"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv *hanlder.NewHelloService) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
