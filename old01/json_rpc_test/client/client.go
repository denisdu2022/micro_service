package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	//1.建立连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("连接失败!")
	}

	var reply string //string是具有默认值的
	//这里使用了json,上边的拨号就需要改为net.Dial(),因为rpc.Dial()是go语言默认的Gob编码的拨号方式
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "libai", &reply)
	if err != nil {
		panic("调用失败!")
	}

	fmt.Println(reply)

}
