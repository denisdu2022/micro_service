package main

import "fmt"

//结合03-design测试

func main() {

	//1.初始化
	myClient := InitClient("127.0.0.1:8800")

	//定义传出参数
	var resp string
	//调用HelloWorld函数
	err := myClient.HelloWorld("杜甫", &resp)
	if err != nil {
		fmt.Println("HelloWorld err !")
		return
	}

	//打印resp
	fmt.Println("resp>>> ", resp, "error>>> ", err)
}
