package reponse

import (
	"fmt"
	"time"
)

//用户返回信息
//自定义个一个类型

type JsonTime time.Time

//转换方法

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	//用户Id
	Id int32 `json:"id"`
	//用户名
	NickName string `json:"name"`
	//生日
	//BirthDay string `json:"birthday"`
	//使用time类型,会自动生成time的json格式
	//BirthDay time.Time `json:"birthday"`

	//推荐使用
	Birthday JsonTime `json:"birthday"`
	//性别
	Gender string `json:"gender"`
	//手机号
	Mobile string `json:"mobile"`
}
