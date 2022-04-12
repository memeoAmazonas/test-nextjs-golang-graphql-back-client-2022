package model

import "time"

type Post struct {
	UserId   int       `json:"userId"`
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	User     *User     `json:"user"`
	Comments int       `json:"comments"`
	Date     time.Time `json:"date"`
	Likes    int       `json:"likes"`
}

type Comment struct {
	PostId int    `json:"postId" bson:"postId"`
	Name   string `json:"name" bson:"name"`
	Body   string `json:"body" bson:"body"`
	Email  string `json:"email" bson:"email"`
}
type NewPost struct {
	UserID   int    `json:"userId"`
	NameUser string `json:"nameUser"`
	Body     string `json:"body"`
}
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Like struct {
	UserId int `json:"userId"`
	PostId int `json:"postId"`
}
