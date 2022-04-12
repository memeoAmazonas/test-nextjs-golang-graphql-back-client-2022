package client

import (
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func GetPost() ([]*model.Post, error) {

	var response []*model.Post
	err := fetch(fmt.Sprintf("%s/posts", os.Getenv("URL_CLIENT_EXTERNAL")), "posts list", &response)
	// TODO mejorar, colocar en distintos hilos la busqueda a bd
	if err == nil {
		results, err := findPost()
		if err == nil {
			response = append(results, response...)
		}
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

	return response, err
}

func GetCommentByPost(id string) ([]*model.Comment, error) {
	var response []*model.Comment
	err := fetch(fmt.Sprintf("%s/posts/%s/comments", os.Getenv("URL_CLIENT_EXTERNAL"), id), "comments by post", &response)
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
	err := fetch(fmt.Sprintf("%s/users", os.Getenv("URL_CLIENT_EXTERNAL")), "users ", &response)
	if err != nil {
		return nil, err
	}
	locals, err := getUsersLocal()
	if err == nil {
		response = append(response, locals...)
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
