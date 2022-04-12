package main

import (
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/controller"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"log"
	"net/http"
)

const PORT = "3200"

func main() {

	database.GetConnection()

	r := mux.NewRouter()
	r.HandleFunc("/post", controller.GetPost).Methods("GET")
	r.HandleFunc("/post", controller.CreatePost).Methods("POST")
	r.HandleFunc("/post/{postId}/comments", controller.GetCommentByPost).Methods("GET")

	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{email}", controller.FindUser).Methods("GET")

	r.HandleFunc("/comment", controller.CreateComment).Methods("POST")
	r.HandleFunc("/comment/{postId}", controller.FindComment).Methods("GET")

	r.HandleFunc("/like", controller.CreateLike).Methods("POST")
	log.Printf("client connect to http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
