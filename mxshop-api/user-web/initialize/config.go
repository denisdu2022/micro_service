package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

//配置文件隔离
//获取环境变量

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

//初始化配置信息

func InitConfig() {
	//环境变量需要生效,需要重启goland
	//fmt.Println(GetEnvInfo("MXSHOP_DEBUG"))
	debug := GetEnvInfo("MXSHOP_DEBUG")
	//设置前缀
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	//TODO 由于MAC环境下暂未取到自定义环境变量,先判断false
	if !debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}

	//获取viper对象
	v := viper.New()
	//设置文件路径 : 目录的相对路径
	v.SetConfigFile(configFileName)
	//读取文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//这个对象在其他文件中使用 -- 全局变量
	//serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	//记录日志
	zap.S().Infof("配置信息: &v", global.ServerConfig)

	//Viper的功能-动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("配置信息: &v", global.ServerConfig)

	})

}
