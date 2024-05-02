package model

import (
	"gorm.io/gorm"
	"time"
)

//公共的模型

type BaseModel struct {
	//主键id
	ID int32 `gorm:"primaryKey"`
	//创建时间     --指定字段名称column
	CreatedAt time.Time `gorm:"column:add_time"`
	//更新时间
	UpdatedAt time.Time `gorm:"column:update_time"`
	//删除时间  --软删除
	DeletedAt gorm.DeletedAt
	//是否删除,bool类型,默认不写生成的字段名也是is_deleted
	IsDeleted bool
}

//用户表

/* 用户密码的保存问题
1.密文保存密码
2.密文不可反解
3. 对称加密,可以反解.采用非对称加密,不可反解
4. md5 信息摘要算法
密码如果不可反解,在找回密码时:用户通过连接来修改密码
*/

type User struct {
	//继承BaseModel
	BaseModel
	//手机号
	Mobile string `gorm:"index:idx_mobile;unique;type:varchar(11);not null;"`
	//密码  --需要对密码进行加密,所以长度可以设置长一些:varchar(100)
	PassWord string `gorm:"type:varchar(100);not null"`
	//昵称
	NickName string `gorm:"type:varchar(20)"`
	//生日 --需要采用指针的模式,来更好的处理日期时间
	BirthDay *time.Time `gorm:"type:datetime"`
	//性别 --这里使用字符串类型,不使用0或1,是为了在查询时可读性高,并没有0或1的性能高
	Gender string `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女,male表示男,默认值为male'"`
	//Role 来区分普通用户和管理员用户
	Role int `gorm:"column:role;default:1;type:int comment '1表示普通用户,2表示管理员'"`
}
