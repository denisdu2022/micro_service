package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//Server端
//-------------------------------------------------------------------
//要求服务端在注册RPC对象时,能让编译器去检测出 注册对象是否是合法对象

//1.创建接口,在接口中去定义方法的原型(不需要实现)

type MyInterface interface {
	HelloWorld(string, *string) error
}

//2.调用方法,调用该方法时,需要给i传参,传参的参数应该是实现了HelloWorld方法的类对象!

func RegisterServer(i MyInterface) {
	rpc.RegisterName("hello", i)
}

//Client端
//-------------------------------------------------------------------
//向调用本地函数一样,去调用远程函数

//定义一个类

type MyClient struct {
	//结构体成员作为类使用,所以成员使用指针,这样就变成了引用传递
	c *rpc.Client
}

//由于使用了c调用了Call,因此需要初始化c

func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	//函数返回结构体对象
	return MyClient{
		c: conn,
	}
}

//实现类的函数,原型参照上边的interface{} 来实现

func (m *MyClient) HelloWorld(a string, b *string) error {
	//参数1应该参照上面的interface{} ,RegisterName而来,a:传入参数  b:传出参数
	//Call方法的返回值本身就是一个error,所以可以直接return
	return m.c.Call("hello.HelloWorld", a, b)
}
