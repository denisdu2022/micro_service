package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

//rpc远程过程调用,如何做到想本地调用

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//使用第三方库
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//fmt.Println("服务端返回的数据: ", string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	fmt.Println(Add(3, 2))
}
