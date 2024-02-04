package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

// rpc远程过程调用,像本地调用一样

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//使用第三方库做测试
	//传输协议:http
	req := HttpRequest.NewRequest()
	//发起get请求
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//打印结果
	//fmt.Println(string(body))
	//反序列化
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	fmt.Println("Sum Value: ", Add(3, 4))
}
