package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	//用户路由分组,改善代码结构性
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关的URL")
	{
		//查询用户列表
		UserRouter.GET("list", api.GetUserList)
	}

}
