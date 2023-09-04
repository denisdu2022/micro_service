package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	//http://127.0.0.1:8000/add?a=1&b=2
	//返回的格式化: json {"data":3}
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()                 //解析参数
		fmt.Println("path: ", request.URL.Path) //callID  request.URL.Path ,数据传输协议 http协议

		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		//设置header
		writer.Header().Set("Content-Type", "application/json")
		//序列化
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		//返回
		_, _ = writer.Write(jData)
	})

	http.ListenAndServe(":8000", nil)

}
