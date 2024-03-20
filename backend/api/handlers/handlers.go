package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/sns/backend/models"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	println(page)

	article := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	println(articleID)

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Comment1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
