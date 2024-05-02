package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

// 定义全局DB对象

var (
	DB *gorm.DB
)

//在go语言中,任何一个package中定义了init方法,方法会自动执行,所以不用显示的去执行

func init() {
	//数据库初始化
	//数据库连接信息
	dsn := "root:password@tcp(hostip)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			//SlowThreshold: time.Second, //慢SQL阈值
			LogLevel: logger.Info, //Log lever(日志级别)
		},
	)
	var err error

	//连接打开数据库 gorm.Open连接数据库  &gorm.Config{}是配置使用
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志配置
		Logger: newLogger,
		//默认创建的表名加s,比如user 创建后为users
		//这里使用NamingStrategy 里配置 schema.NamingStrategy 的SingularTable 为true 就不会加s了
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败!")
	}
	//db对象类似文件句柄,后续对数据库操作都使用db对象
	//fmt.Println(DB)

	////迁移模型类: 将模型类转换为SQL表
	////db.AutoMigrate(&Teacher{})
	//err = DB.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err.Error())
	//}
}
