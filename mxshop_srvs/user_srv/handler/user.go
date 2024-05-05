package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"time"
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
		Mobile:   user.Mobile,
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

//通过手机号查询用户信息

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	//数据库查询
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	//如果返回的行数等于0,则没有查询到用户
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	//判断是否存在其他的错误,比如数据库连接不上等
	if result.Error != nil {
		return nil, result.Error
	}

	//填充返回用户信息
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

//通过ID查询用户

func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {

	var user model.User
	//由于ID是主键,查询比较简单
	result := global.DB.Find(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

//创建用户

func (s *UserServer) CreteUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//先查询用户是否存在
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	//用户已存在,则不能添加
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	//数据库中存储加密算法,盐值和encodedPwd
	user.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	//将用户信息存储到数据库中
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	//用户信息保存到数据库后,用户ID需要返回给调用方
	userInfoRsp := ModelToResponse(user)

	return &userInfoRsp, nil
}

//个人中心更新用户

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	//在更新用户时,用户肯定是存在的,所以需要先查询用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	//生日的处理,转换
	birthDay := time.Unix(int64(req.BirthDay), 0)

	//更新用户信息
	user.NickName = req.NickName
	user.BirthDay = &birthDay
	user.Gender = req.Gender

	//保存到数据库中
	result = global.DB.Save(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return &empty.Empty{}, nil
}

//检查用户密码

func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {

	//解析
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")

	//验证 password.Verify(原始密码,salt,encodedPwd,options)
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check)

	return &proto.CheckResponse{Success: check}, nil
}
