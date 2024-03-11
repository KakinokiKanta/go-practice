package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/gorilla/mux"
)

// POST: ブログ記事の投稿をするためのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// GET: ブログ記事の一覧を取得するためのハンドラ
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

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// GET: 指定した記事ナンバーの投稿データを取得するためのハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// POST: 記事にいいねをつけるためのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// POST: 記事にコメントを投稿するためのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
