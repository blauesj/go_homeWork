package main

import (
	"fmt"
	"go_homeWork/chapter3/question_5/models"
)

func main() {
	models.Init()

	user, _ := models.GetUser(1)
	fmt.Printf("%+v\n", user)

	posts, _ := models.GetMostCommentedPosts()

	fmt.Printf("%+v\n", posts)

	//models.SavePost(&models.Post{UserId: 1, Title: "Hello World", Content: "Hello World"})

	//models.AddComment(&models.Comment{UserId: 1, PostId: 9, Content: "Hello World222"})
	models.DelComment(6)
}
