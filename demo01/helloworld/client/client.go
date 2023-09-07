package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	//1.建立连接
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("连接服务端失败>>> ", err)
	}

	//2.想调用本地函数一样调用远程函数
	//定义指针变量
	var reply *string = new(string)
	//fmt.Println(*reply)
	//服务名.方法名 传入参数  传出参数（由于是指针变量，需要&reply写入,如果使用reply 会报错panic: runtime error: invalid memory address or nil pointer dereference ）
	//err = client.Call("HelloService.Hello", "李白", &reply)
	//或者使用new先开辟内存空间，在把值赋值给指针变量
	err = client.Call("HelloService.Hello", "李白", reply)
	if err != nil {
		panic("调用失败...")
	}

	fmt.Println("reply>>> ", *reply)
}
