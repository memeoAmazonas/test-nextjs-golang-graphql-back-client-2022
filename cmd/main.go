package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/controller"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/database"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-client-2022/internal/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const defaultPort = "3200"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	var httpAddr = flag.String("http", ":"+defaultPort, "HTTP listen address")
	flag.Parse()
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		if err := database.GetConnection(); err != nil {
			errs <- err
		}
	}()

	r := mux.NewRouter()
	r.Use(middleware.HeaderMiddleware)
	r.HandleFunc("/post", controller.GetPost).Methods("GET")
	r.HandleFunc("/post/{postId}/comments", controller.GetCommentByPost).Methods("GET")
	r.HandleFunc("/user/{email}", controller.FindUser).Methods("GET")
	r.HandleFunc("/comment/{postId}", controller.FindComment).Methods("GET")

	r.HandleFunc("/post", controller.CreatePost).Methods("POST")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/comment", controller.CreateComment).Methods("POST")
	r.HandleFunc("/like", controller.CreateLike).Methods("POST")

	go func() {
		handler := cors.Default().Handler(r)
		log.Print("listen on port ", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	log.Fatal("exit", <-errs)

}
