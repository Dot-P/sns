package main

import (
	"log"
	"net/http"

	"github.com/sns/backend/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHander)
	http.HandleFunc("/article/1", handlers.ArticleDetailHander)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHander)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
