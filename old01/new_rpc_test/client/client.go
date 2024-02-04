package main

import (
	"fmt"
	"new_rpc/client_proxy"
)

func main() {
	//1.启动连接

	client := client_proxy.NewHelloServiceClient("tcp", "localhost:8080")

	var reply string
	//2. 调用服务
	err := client.Hello("李白", &reply)
	if err != nil {
		panic("调用失败!")
	}

	fmt.Println(reply)
}
