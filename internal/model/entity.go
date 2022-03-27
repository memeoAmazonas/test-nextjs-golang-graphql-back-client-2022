package model

type Post struct {
	UserId   int    `json:"userId"`
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	User     *User  `json:"user"`
	Comments int    `json:"comments"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Name   string `json:"name"`
	Body   string `json:"body"`
	Email  string `json:"email"`
}
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
