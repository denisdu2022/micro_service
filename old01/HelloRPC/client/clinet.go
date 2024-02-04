package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	//1.建立连接
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		panic("连接失败!")
	}

	var reply string
	//2. 调用服务
	err = client.Call("HelloService.Hello", "libai", &reply)
	if err != nil {
		panic("调用失败!")
	}
	fmt.Println(reply)
}
