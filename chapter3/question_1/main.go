package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//stu := Student{
	//	Id:    1,
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//db.Create(&stu)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Student
	db.Where("age > ?", 18).Find(&students)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。

	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

	//db.Where("age < ?", 15).Delete(&Student{})
}

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func (Student) TableName() string {
	return "student"
}
