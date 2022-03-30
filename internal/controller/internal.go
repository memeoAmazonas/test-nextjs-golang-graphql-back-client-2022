package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/client"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/mapper"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CreateComment(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comment *model.Comment
	log.Info("Create comment")
	if err := json.NewDecoder(request.Body).Decode(&comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Create comment ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	_, err := client.CreateComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create comment ", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	log.Info("Create comment successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("success")
}

func FindComment(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get comments")
	vars := mux.Vars(request)
	id := vars["postId"]
	toSend, err := strconv.Atoi(id)
	if err != nil {
		log.Error("Get comments ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	results, err := client.FindComment(toSend)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Get comments ", err)
		json.NewEncoder(w).Encode(err)
	}

	log.Info("Get comments succesfully")

	json.NewEncoder(w).Encode(results)
}

func CreatePost(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post *model.NewPost
	log.Info("Create post")
	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Create post ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	_, err := client.CreatePost(mapper.NewPostMapper(post))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create post ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Info("Create post successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("success")
}
func CreateUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *model.User
	log.Info("Create user")
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Create user ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	id, err := client.CreateUser(user)
	fmt.Println("id", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create user ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Info("Create user successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}
