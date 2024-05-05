package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
)

func main() {
	//port := 8021

	//1.初始化logger
	initialize.InitLogger()

	//初始化配置文件
	initialize.InitConfig()

	//2.初始化Routers
	Router := initialize.Routers()
	//3.启动
	zap.S().Infof("启动服务器,端口: %d", global.ServerConfig.Port)
	//在启动时,可能会有错误,比如端口被占用,所以这里顺便处理错误
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}

}
