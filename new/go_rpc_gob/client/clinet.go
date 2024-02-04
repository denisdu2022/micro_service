package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.建立连接
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		panic("建立连接失败!")
	}
	var reply string //string有默认值,空字符串
	err = client.Call("HelloService.Hello", "李白", &reply)
	if err != nil {
		panic("调用失败!")
	}

	fmt.Println("reply>>> ", reply)
}
