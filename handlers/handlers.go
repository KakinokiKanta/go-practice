package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/KakinokiKanta/go-intermediate/services"
	"github.com/gorilla/mux"
)

/*
POST /article
ブログ記事の投稿をするためのハンドラ
*/
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

/*
GET /article/list
ブログ記事の一覧を取得するためのハンドラ
*/
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータを読み出し
	queryMap := req.URL.Query()

	// レスポンスで返す一覧のページ番号
	var page int

	// クエリパラメータに"page"があれば
	// その値のページ番号を変数pageに格納
	// なければ1をpageに格納
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

	articleList, err := services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

/*
GET /article/{id}
指定した記事ナンバーの投稿データを取得するためのハンドラ
*/
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// muxを用いてリクエストからidを抽出
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

/*
POST /article/nice
記事にいいねをつけるためのハンドラ
*/
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

/*
POST /comment
記事にコメントを投稿するためのハンドラ
*/
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
