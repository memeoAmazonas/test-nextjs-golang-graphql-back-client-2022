package mapper

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	"math/rand"
)

func NewPostMapper(post *model.NewPost) *model.Post {
	title := post.Body
	if len(post.Body) > 10 {
		title = post.Body[:10]
	}
	return &model.Post{
		UserId: post.UserID,
		Id:     rand.Intn(2000000),
		Title:  title,
		Body:   post.Body,
		User: &model.User{
			Name: post.NameUser,
			Id:   post.UserID,
		},
	}
}
func ReverseListPost(list []*model.Post) []*model.Post {
	var response []*model.Post
	for i := len(list) - 1; i >= 0; i-- {
		response = append(response, list[i])
	}
	return response
}
