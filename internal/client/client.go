package client

import (
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetPost() ([]*model.Post, error) {
	var response []*model.Post
	log.Info("get post client")
	res, err := http.Get("http://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Error("Get post call", err.Error())
		return nil, err
	}
	log.Info("success call get post")
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse data get post", err.Error())
		return nil, err
	}
	log.Info("successfull get post")
	return response, nil
}
func GetCommentByPost(id string) ([]*model.Comment, error) {
	var response []*model.Comment
	log.Info("get comments by post client")
	res, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%s/comments", id))

	if err != nil {
		log.Error("Get comments by post call", err.Error())
		return nil, err
	}
	log.Info("success call get comments by post")

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse data comments by post", err.Error())
		return nil, err
	}
	return response, nil

}
