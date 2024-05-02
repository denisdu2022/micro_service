package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

//MD5的方法

func genMD5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	////数据库初始化
	////数据库连接信息
	//dsn := dsn := "root:password@tcp(hostip)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//
	////创建日志对象
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
	//	logger.Config{
	//		//SlowThreshold: time.Second, //慢SQL阈值
	//		LogLevel: logger.Info, //Log lever(日志级别)
	//	},
	//)
	//
	////连接打开数据库 gorm.Open连接数据库  &gorm.Config{}是配置使用
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	//日志配置
	//	Logger: newLogger,
	//	//默认创建的表名加s,比如user 创建后为users
	//	//这里使用NamingStrategy 里配置 schema.NamingStrategy 的SingularTable 为true 就不会加s了
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//})
	//if err != nil {
	//	panic("数据库连接失败!")
	//}
	////db对象类似文件句柄,后续对数据库操作都使用db对象
	//fmt.Println(db)
	//
	////迁移模型类: 将模型类转换为SQL表
	////db.AutoMigrate(&Teacher{})
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err.Error())
	//}

	//生成MD5
	//fmt.Println(genMD5("123456")) //e10adc3949ba59abbe56e057f20f883e

	//// go-password-encoder 的基本用法
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	rawPwd := "AdM^Rt!e$s5j"
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(rawPwd, options)
	//数据库中存储加密算法,盐值和encodedPwd
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//fmt.Println(len(salt + encodedPwd))
	fmt.Println(newPassword)
	fmt.Println(len(newPassword))

	//解析
	passwordInfo := strings.Split(newPassword, "$")
	for i, v := range passwordInfo {
		fmt.Println(i, v)
		/*
			0
			1 pbkdf2-sha512
			2 67H93DQAJlJlXzIV
			3 06feadd1ecf5268dcdf0255dca5c128927538852b114dd0fea451a6e16534309
		*/
	}

	//验证 password.Verify(原始密码,salt,encodedPwd,options)
	check := password.Verify(rawPwd, passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true

}
