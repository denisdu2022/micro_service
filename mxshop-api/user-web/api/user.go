package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/reponse"
	"mxshop-api/user-web/proto"
	"net/http"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code 转换为http的状态码
	if err != nil {
		//status是grpc的
		if e, ok := status.FromError(err); ok {
			//e.code是状态码 ,e.message是返回信息
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					//e.Message() 开发中使用
					"msg": "其他错误" + e.Message(),
				})
			}
		}
		return
	}
}

//获取用户列表

func GetUserList(ctx *gin.Context) {
	//ip := "127.0.0.1"
	//port := 50051
	//拨号连接用户grpc服务
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg:", err.Error())
	}

	//生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	//获取用户列表
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败")
		//将GRPC的状态码转换为http的状态码
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	//返回信息 :有多个用户,可以是map,也可以是[]slice
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//使用map也没有什么问题
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["name"] = value.NickName
		//data["birthday"] = value.BirthDay
		//data["gender"] = value.Gender
		//data["mobile"] = value.Mobile

		//推荐使用
		user := reponse.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//需要做转换
			//转换1:当是BirthDay string `json:"birthday"` 类型时
			//BirthDay: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			//BirthDay: time.Time(time.Unix(int64(value.BirthDay), 0)),
			//推荐使用
			Birthday: reponse.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}

		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}
