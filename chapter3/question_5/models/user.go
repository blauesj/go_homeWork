package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"colum:user_name;type:varchar(50);size:50"`
	Password  string `gorm:"colum:password;type:varchar(50);size:50"`
	PostCount uint32 `gorm:"colum:post_count;type:int(11);default:0;not null"`
	Posts     []Post
}

func GetUser(id uint) (User, error) {
	var user User
	err := DB.Model(&User{}).Preload("Posts").Preload("Posts.Comments").Find(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
