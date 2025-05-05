package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	sdn := "root:mysql@tcp(127.0.0.1:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建数据库连接
	db, err := gorm.Open(mysql.Open(sdn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	// 创建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	DB = db
}
