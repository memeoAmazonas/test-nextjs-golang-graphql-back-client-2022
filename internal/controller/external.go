package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/client"
	"net/http"
)

func GetPost(w http.ResponseWriter, _ *http.Request) {
	post, err := client.GetPost()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(post)
}

func GetCommentByPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["postId"]
	comments, err := client.GetCommentByPost(id)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(comments)
}
