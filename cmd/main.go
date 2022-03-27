package main

import (
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/controller"
	"log"
	"net/http"
)

const PORT = "3002"

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/post", controller.GetPost).Methods("GET")
	r.HandleFunc("/post/{postId}/comments", controller.GetCommentByPost).Methods("GET")
	log.Printf("client connect to http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
