package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructture:"port"`
}

//映射成struct

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

//配置文件隔离
//获取环境变量

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {

	//环境变量需要生效,需要重启goland
	//fmt.Println(GetEnvInfo("MXSHOP_DEBUG"))
	debug := GetEnvInfo("MXSHOP_DEBUG")
	//设置前缀
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("viper_test/ch02/%s-pro.yaml", configFilePrefix)
	//TODO 由于MAC环境下暂未取到自定义环境变量,先判断false
	if !debug {
		configFileName = fmt.Sprintf("viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}

	//获取viper对象
	v := viper.New()
	//设置文件路径 : 目录的相对路径
	v.SetConfigFile(configFileName)
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

	//Viper的功能-动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 300)

}
