package controller

import (
	"encoding/json"
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
	var post *model.NewPost
	log.Info("Create post")
	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Create post ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	result, err := client.CreatePost(mapper.NewPostMapper(post))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create post ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Info("Create post successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
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
	send, err := client.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create user ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Info("Create user successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(send)
}

func FindUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get user")
	vars := mux.Vars(request)
	email := vars["email"]
	if email == "" {
		log.Error("Get user ", "email required")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email required")
	}
	result, err := client.GetUser(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Get user ", err)
		json.NewEncoder(w).Encode(err)
	}

	log.Info("Get user succesfully")

	json.NewEncoder(w).Encode(result)
}

func CreateLike(w http.ResponseWriter, request *http.Request) {
	var like *model.Like
	log.Info("Create like")
	if err := json.NewDecoder(request.Body).Decode(&like); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Create like ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	created, err := client.CreateLike(like)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Create like ", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := client.UpdatePostLike(like.PostId, created); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Update like ", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	response := "deleted"
	if created {
		response = "created"
	}
	log.Info(response + " like successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
