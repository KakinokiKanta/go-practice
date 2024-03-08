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
	r.HandleFunc("/hello", handlers.HelloHandler)
	// ブログ記事の投稿をするためのエンドポイント
	r.HandleFunc("/article", handlers.PostArticleHandler)
	// ブログ記事の一覧を取得するためのエンドポイント
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// 記事にいいねをつけるためのエンドポイント
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// 記事にコメントを投稿するためのエンドポイント
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
