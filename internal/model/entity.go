package model

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Name   string `json:"name"`
	Body   string `json:"body"`
}
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
