package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1.建立连接
	conn, err := net.Dial("tcp", "localhost:8000")
	//错误处理
	if err != nil {
		panic("建立连接失败!")
	}
	//2.调用
	var reply string
	//获取客户端,使用json编码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	_ = client.Call("HelloService.Hello", "李白", &reply)
	//错误处理
	if err != nil {
		panic("调用失败!")
	}
	fmt.Println("reply: ", reply)
}
