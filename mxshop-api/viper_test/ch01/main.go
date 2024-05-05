package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//映射成struct

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

func main() {

	//获取viper对象
	v := viper.New()
	//设置文件路径 : 目录的相对路径
	v.SetConfigFile("viper_test/ch01/config.yaml")
	//读取文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	//获取数据
	//fmt.Println(v.Get("name"))
	fmt.Println(serverConfig)
}
