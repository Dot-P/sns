package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func ArticleListHander(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List")
}
func ArticleDetailHander(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d", articleID)
	io.WriteString(w, resString)
}
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...")
}
func PostCommentHander(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...")
}
