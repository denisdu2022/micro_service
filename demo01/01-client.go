package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {

	//1.用rpc拨号 连接服务器
	//"127.0.0.1:8800" 是远程服务器的地址和端口
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")

	//使用jsonRpc
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")

	if err != nil {
		fmt.Println("rpc.Dial err >>> ", err)
		return
	}

	defer conn.Close()

	//2.调用远程函数
	var reply string //接收函数返回值
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call err >>> ", err)
		return
	}

	//3.打印函数返回值
	fmt.Println("reply>>> ", reply)
}
