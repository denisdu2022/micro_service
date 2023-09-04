package main

import (
	"fmt"
	"micro_service/test/old_terst01/calc"
)

//1.一定要将代码新建到GOPATH下的src下面
//2.go env -w GO111MODULE=off
func main() {
	fmt.Println("hello!")
	rep := calc.Add(3, 4)
	fmt.Println("rep > > > ", rep)

}
