package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main () {
	// helloHandlerという名前で、ハンドラを定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理内容:
		// 何がきても、"Hello, world!"の文字列を返す
		io.WriteString(w, "Hello, world!\n")
	}

	// TODO: ブログ記事の投稿をするためのハンドラ
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	// TODO: ブログ記事の一覧を取得するためのハンドラ
	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	// TODO: 指定した記事ナンバーの投稿データを取得するためのハンドラ
	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	}

	// TODO: 記事にいいねをつけるためのハンドラ
	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	// TODO: 記事にコメントを投稿するためのハンドラ
	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...")
	}

	// 定義したhelloHandlerを使うように登録
	http.HandleFunc("/hello", helloHandler)
	// ブログ記事の投稿をするためのエンドポイント
	http.HandleFunc("/article", postArticleHandler)
	// ブログ記事の一覧を取得するためのエンドポイント
	http.HandleFunc("/article/list", articleListHandler)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	http.HandleFunc("/article/1", articleDetailHandler)
	// 記事にいいねをつけるためのエンドポイント
	http.HandleFunc("/article/nice", postNiceHandler)
	// 記事にコメントを投稿するためのエンドポイント
	http.HandleFunc("/comment", postCommentHandler)

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
