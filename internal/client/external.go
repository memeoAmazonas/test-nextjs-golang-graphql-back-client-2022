package client

import (
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetPost() ([]*model.Post, error) {

	var response []*model.Post
	err := fetch("http://jsonplaceholder.typicode.com/posts", "posts list", &response)
	if err == nil {
		users, err := getUsers()
		if err == nil {
			for _, it := range response {
				act := users[it.UserId]
				if act != nil {
					com, err := GetCommentByPost(strconv.Itoa(it.Id))
					if err == nil {
						it.Comments = len(com)
					}
					it.User = act
				} else {

					it.User = &model.User{
						Id:    0,
						Name:  "anonymous",
						Email: "anonymous@anonymous.com",
					}
				}
			}
		}
	}
	results, err := FindPost(response)
	if err != nil {
		log.Error("Add local post")
	}
	return results, err
}
func GetCommentByPost(id string) ([]*model.Comment, error) {
	var response []*model.Comment
	err := fetch(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%s/comments", id), "comments by post", &response)
	send, _ := strconv.Atoi(id)
	locales, err := FindComment(send)
	if err == nil {
		for _, it := range locales {
			response = append(response, it)
		}
	}
	return response, err
}
func getUsers() (map[int]*model.User, error) {
	users := make(map[int]*model.User)
	var response []*model.User
	err := fetch("http://jsonplaceholder.typicode.com/users", "users ", &response)
	if err != nil {
		return users, err
	}
	for _, it := range response {
		if users[it.Id] == nil {
			users[it.Id] = it
		}
	}
	return users, nil

}
func fetch(url, logKey string, decode interface{}) error {
	log.Info("get " + logKey)
	res, err := http.Get(url)
	if err != nil {
		log.Error("Get "+logKey, err.Error())
		return err
	}
	log.Info("success call " + logKey)
	if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {
		log.Error("Parse data "+logKey, err.Error())
		return err
	}
	log.Info("successfull get " + logKey)
	return nil
}
