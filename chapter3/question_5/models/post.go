package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserId        uint   `gorm:"colum:user_id;type:int"`
	Title         string `gorm:"colum:title;type:varchar(50)"`
	Content       string `gorm:"colum:content;type:text"`
	CommentsCount int    `gorm:"colum:comments_count;type:int"`
	CommentStatus string `gorm:"colum:comment_status;type:varchar(20);default:无评论"`
	Comments      []Comment
}

func GetMostCommentedPosts() (Post, error) {
	var post Post
	err := DB.Preload("Comments").Order("comments_count desc").Limit(1).Find(&post).Error
	return post, err
}

func SavePost(post *Post) error {
	return DB.Create(post).Error
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", post.UserId).Update("post_count", gorm.Expr("post_count + ?", 1))
	return nil
}
