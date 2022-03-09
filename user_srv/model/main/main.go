package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/Dlimingliang/shop_srvs/user_srv/model"
)

func main() {
	//日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	//建立数据库连接
	dsn := "root:123456!@tcp(127.0.0.1:3306)/shop_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //配置单数表名,默认如果struct为User,生成的表为users.配置单数表名生成的表为user
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//生成表结构
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
