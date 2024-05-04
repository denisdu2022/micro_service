package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//获取用户列表

func GetUserList(ctx *gin.Context) {

	zap.S().Debug("获取用户列表页")
}
