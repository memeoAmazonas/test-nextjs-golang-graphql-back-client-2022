package main

import (
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/controller"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/post", controller.GetPost).Methods("GET")
	r.HandleFunc("/post/{postId}/comments", controller.GetCommentByPost).Methods("GET")
	log.Fatal(http.ListenAndServe(":3002", r))
}
