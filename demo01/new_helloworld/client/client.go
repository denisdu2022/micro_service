package main

import (
	"fmt"
	"newhelloworld/client_proxy"
)

func main() {
	//1.建立连接
	//client, err := rpc.Dial("tcp", "localhost:8800")
	////错误处理
	//if err != nil {
	//	panic("连接失败!")
	//}
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:8800")

	//2.向Server发送数据
	var reply string

	//err = client.Call(hanlder.HelloServiceName+".Hello", "李白", &reply)

	err := client.Hello(
		"李白",
		&reply,
	)
	//错误处理
	if err != nil {
		panic(err)
	}

	fmt.Println("reply>>> ", reply)
}
