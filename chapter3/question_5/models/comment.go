package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Comment struct {
	gorm.Model
	UserId  uint   `gorm:"colum:user_id;type:int"`
	PostId  uint   `gorm:"colum:post_id;type:int"`
	Content string `gorm:"colum:content;type:text"`
}

func (comment *Comment) AfterDelete(tx *gorm.DB) error {
	if tx.RowsAffected == 0 {
		log.Println("delete failed")
		return nil
	}
	var post Post
	tx.Model(&post).Clauses(clause.Returning{}).Where("id = ?", comment.PostId).Update("comments_count", gorm.Expr("comments_count - 1"))
	tx.First(&post, comment.PostId)
	if post.CommentsCount == 0 {
		tx.Model(&post).Where("id = ?", comment.PostId).Update("comment_status", "无评论")
	} else {
		tx.Model(&post).Where("id = ?", comment.PostId).Update("comment_status", "已评论")
	}
	return nil
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	var post Post
	tx.Model(&post).Clauses(clause.Returning{}).Where("id = ?", comment.PostId).Update("comments_count", gorm.Expr("comments_count + 1"))
	tx.First(&post, comment.PostId)
	if post.CommentsCount == 0 {
		tx.Model(&post).Where("id = ?", comment.PostId).Update("comment_status", "无评论")
	} else {
		tx.Model(&post).Where("id = ?", comment.PostId).Update("comment_status", "已评论")
	}
	return nil
}

func AddComment(comment *Comment) error {
	err := DB.Create(&comment).Error
	return err
}

func DelComment(id int) error {
	err := DB.Delete(&Comment{}, id).Error

	return err
}
