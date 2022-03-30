package client

import (
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/mapper"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
)

const (
	COMMENT = "comments"
	POST    = "posts"
	USER    = "user"
	DB      = "test_marzo_go_nextjs_grahql"
)

func CreateComment(comment *model.Comment) (interface{}, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, DB, COMMENT, comment)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cursor.InsertedID, nil
}

func FindComment(id int) ([]*model.Comment, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	cursor, err := database.Query(client, ctx, DB, COMMENT, bson.D{{"postId", id}})
	if err != nil {
		return nil, err
	}
	var results []*model.Comment
	for cursor.Next(ctx) {
		var v *model.Comment
		err := cursor.Decode(&v)
		if err != nil {
			log.Error(err.Error())
		}
		results = append(results, v)
	}
	return results, nil
}
func FindPost(res []*model.Post) ([]*model.Post, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	cursor, err := database.Query(client, ctx, DB, POST, bson.D{})
	if err != nil {
		return nil, err
	}
	results := res
	for cursor.Next(ctx) {
		var v *model.Post
		err := cursor.Decode(&v)
		if err != nil {
			log.Error(err.Error())
		}
		results = append(results, v)
	}
	val := mapper.ReverseListPost(results)
	return val, nil
}
func CreatePost(post *model.Post) (interface{}, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer database.Close(client, ctx, cancel)
	_, err = database.SaveOne(client, ctx, DB, POST, post)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return post.Id, nil
}
func CreateUser(user *model.User) (int, error) {
	uuid := rand.Intn(2000000)
	user.Id = uuid
	user.Email = fmt.Sprintf("%s@tucan.com", user.Name)
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
		return -1, err
	}
	defer database.Close(client, ctx, cancel)
	_, err = database.SaveOne(client, ctx, DB, USER, user)
	if err != nil {
		log.Error(err.Error())
		return -1, err
	}
	return uuid, nil
}
