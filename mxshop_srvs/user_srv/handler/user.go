package handler

import (
	"context"
	"gorm.io/gorm"
	"user_srv/global"
	"user_srv/model"
	"user_srv/proto"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

//提供一个可以公用的UserInfoResponse方法

func ModelToResponse(user model.User) proto.UserInfoResponse {
	//在grpc的message中的字段有默认值,不能随便赋值nil进去,容易报错
	//这里要清楚,哪些字段有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.PassWord,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}

	if user.BirthDay != nil {
		//字段赋值要注意类型对应
		userInfoRsp.BirthDay = uint64(user.BirthDay.Unix())
	}
	return userInfoRsp
}

//gorm的Scopes中提供的分页方法:https://gorm.io/zh_CN/docs/scopes.html

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//go语言中使用鸭子类型,所以不需要继承具体类型,只需要实现相关方法即可

// 获取用户列表

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表信息
	var users []model.User
	//数据库查询
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	//把返回的总行数赋值给rsp总的大小
	rsp.Total = int32(result.RowsAffected)

	//分页
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	//数据填充到rsp
	for _, user := range users {
		//先赋值变量
		userInfoRsp := ModelToResponse(user)
		//在&取地址,这样做是内存安全的,符合go的语法,否则会报错
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}
