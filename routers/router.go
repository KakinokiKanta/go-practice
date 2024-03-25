package routers

import (
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()

	// ブログ記事の投稿をするためのエンドポイント
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	// ブログ記事の一覧を取得するためのエンドポイント
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	// 記事にいいねをつけるためのエンドポイント
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)

	// 記事にコメントを投稿するためのエンドポイント
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	return r
}
