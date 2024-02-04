package main

import (
	"fmt"
	"new_rpc_test/client_proxy"
)

func main() {
	//1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:8000")

	//2.本地调用
	var reply string
	err := client.Hello("李白", &reply)
	if err != nil {
		panic("调用失败!")
	}

	fmt.Println(reply)
}
