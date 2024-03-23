package main

import (
	"log"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/controllers"
	"github.com/KakinokiKanta/go-intermediate/services"
	"github.com/gorilla/mux"
)

func main() {
	ser := services.NewMyAppService([])
	con := controllers.NewMyAppController(ser)
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

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
