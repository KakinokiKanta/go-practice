package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// /helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// TODO: ブログ記事の投稿をするためのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// TODO: ブログ記事の一覧を取得するためのハンドラ
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

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

// TODO: 指定した記事ナンバーの投稿データを取得するためのハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// TODO: 記事にいいねをつけるためのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// TODO: 記事にコメントを投稿するためのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...")
}