package global

import "mxshop-api/user-web/config"

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{} //后边需要改变它,所以必须是指针类型
)
