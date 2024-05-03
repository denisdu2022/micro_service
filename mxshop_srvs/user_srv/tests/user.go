package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"user_srv/proto"
)

//全局的UserClient对象

var userClient proto.UserClient

//全局的ClientConn

var conn *grpc.ClientConn

func Init() {
	//建立RPC连接
	var err error
	conn, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)

}

//创建用户

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreteUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("baijuyi%d", i),
			Mobile:   fmt.Sprintf("9874561230%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

//获取用户信息列表

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 20,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		//校验密码
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		//fmt.Println(user)
	}
}

//通过ID获取用户信息

func TestGetUserInfoById() {
	rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: 4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

//通过手机号获取用户信息

func TestGetUserInfoByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "98745612300"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

//更新用户信息

func TestUpdateUserInfo() {
	rsp, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       4,
		NickName: "menghaoran",
		Gender:   "male",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func main() {
	//初始化rpc连接
	Init()

	//TestCreateUser()
	//TestGetUserList()
	TestGetUserInfoById()
	//TestGetUserInfoByMobile()
	//TestUpdateUserInfo()

	conn.Close()

}
