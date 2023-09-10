package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	//1.建立连接
	//conn,err := rpc.Dial("tcp",":8000")
	//由于使用了json-rpc,所以需要使用net包拨号
	conn, err := net.Dial("tcp", "127.0.0.1:8800")
	//错误处理
	if err != nil {
		fmt.Println("连接失败!")
		panic(err)
	}

	var reply string

	//2.json序列化
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//3.向server端发送数据
	err = client.Call("HelloService.Hello", "李白", &reply)
	//错误处理
	if err != nil {
		panic(err)
	}

	fmt.Println("reply>>> ", reply)
}

//json发送的数据格式
// {"method":"HelloService.Hello","params":["hello"],"id":0}
