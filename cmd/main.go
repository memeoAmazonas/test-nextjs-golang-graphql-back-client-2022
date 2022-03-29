package main

import (
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/controller"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"log"
	"net/http"
)

const PORT = "3002"

func main() {

	database.GetConnection()

	r := mux.NewRouter()
	r.HandleFunc("/post", controller.GetPost).Methods("GET")
	r.HandleFunc("/post/{postId}/comments", controller.GetCommentByPost).Methods("GET")
	r.HandleFunc("/comment", controller.CreateComment).Methods("POST")
	r.HandleFunc("/comment/{postId}", controller.FindComment).Methods("GET")
	log.Printf("client connect to http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
