package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "hello, world!\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func ArticleListHander(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func ArticleDetailHander(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func PostCommentHander(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
