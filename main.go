package main

import (
	"log"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/handlers"
	"github.com/gorilla/mux"
)

func main () {
	r := mux.NewRouter()

	// 定義したhelloHandlerを使うように登録
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	// ブログ記事の投稿をするためのエンドポイント
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	// ブログ記事の一覧を取得するためのエンドポイント
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	// 記事にいいねをつけるためのエンドポイント
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	// 記事にコメントを投稿するためのエンドポイント
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
