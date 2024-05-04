package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/router"
)

func Routers() *gin.Engine {
	//获取gin引擎
	Router := gin.Default()
	//全局的group
	ApiGroup := Router.Group("/u/v1")
	//调用用户路由分组
	router.InitUserRouter(ApiGroup)

	//返回gin引擎
	return Router
}
