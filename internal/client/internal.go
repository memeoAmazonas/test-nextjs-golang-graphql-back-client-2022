package client

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COMMENT = "comments"
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
