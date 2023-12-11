package db

import (
	"Todolist/pkg/constants"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func InitDB() {

	Init(constants.Connect)
}
func Init(connect string) {
	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //开启日志
	})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetMaxIdleConns(10)
}
