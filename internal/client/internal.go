package client

import (
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"time"
)

const (
	COMMENT = "comments"
	POST    = "posts"
	USER    = "user"
	LIKE    = "like"
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
	cursor, err := database.Query(client, ctx, DB, COMMENT, bson.D{{"postId", id}}, nil)
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

func findPost() ([]*model.Post, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	opts := options.Find().SetSort(bson.D{{"date", -1}})
	cursor, err := database.Query(client, ctx, DB, POST, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var results []*model.Post
	for cursor.Next(ctx) {
		var v *model.Post
		err := cursor.Decode(&v)
		if err != nil {
			log.Error(err.Error())
		}
		results = append(results, v)
	}
	return results, nil
}

func CreatePost(post *model.Post) (*model.Post, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	post.Date = time.Now()
	defer database.Close(client, ctx, cancel)
	_, err = database.SaveOne(client, ctx, DB, POST, post)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return post, nil
}

func CreateUser(user *model.User) (int, error) {
	uuid := rand.Intn(2000000)
	user.Id = uuid
	exist, err := GetUser(user.Email)
	if exist != nil {
		log.Error("User exist")
		return -2, err
	}
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

func getUsersLocal() ([]*model.User, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	cursor, err := database.Query(client, ctx, DB, USER, bson.D{}, nil)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for cursor.Next(ctx) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			log.Error(err.Error())
		}
		users = append(users, v)
	}
	return users, nil
}

func GetUser(email string) (*model.User, error) {
	client, ctx, cancel, _ := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	user, err := database.FindOne(client, ctx, DB, USER, bson.D{{"email", email}})
	if err != nil {
		return nil, err
	}
	var v *model.User
	if err := user.Decode(&v); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return v, nil
		}
		log.Error(err)
		return nil, err
	}
	return v, nil
}

func CreateLike(lik *model.Like) (bool, error) {
	client, ctx, cancel, _ := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	like, err := database.FindOne(client, ctx, DB, LIKE, bson.D{{"postid", lik.PostId}})
	if err != nil {
		return false, err
	}
	var v *model.Like
	if err := like.Decode(&v); err != nil {
		if err.Error() == "mongo: no documents in result" {

			if _, er := database.SaveOne(client, ctx, DB, LIKE, lik); er != nil {
				return false, er
			}
			return true, nil
		} else {
			return false, err
		}
	}
	drop, err := database.DeleteOne(client, ctx, DB, LIKE, bson.D{{"postid", lik.PostId}})

	if err != nil {
		return false, err
	}
	fmt.Printf("eliminado %v ", drop.DeletedCount)
	return false, nil

}

func UpdatePostLike(id int, update bool) error {
	client, ctx, cancel, _ := database.Connect("mongodb://localhost:27017")
	defer database.Close(client, ctx, cancel)
	filter := bson.D{{"id", id}}
	toSet := -1
	if update {
		toSet = 1
	}
	upd := bson.D{
		{"$inc", bson.D{
			{"likes", toSet},
		}},
	}
	send, err := database.UpdateOne(client, ctx, DB, POST, filter, upd)
	if err != nil {
		return err
	}
	fmt.Printf("modificado %v  %v \n", send.MatchedCount, send.ModifiedCount)

	return nil
}
