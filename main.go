package main

import (
	"log"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/handlers"
)

func main () {
	// 定義したhelloHandlerを使うように登録
	http.HandleFunc("/hello", handlers.HelloHandler)
	// ブログ記事の投稿をするためのエンドポイント
	http.HandleFunc("/article", handlers.PostArticleHandler)
	// ブログ記事の一覧を取得するためのエンドポイント
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// 記事にいいねをつけるためのエンドポイント
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// 記事にコメントを投稿するためのエンドポイント
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
