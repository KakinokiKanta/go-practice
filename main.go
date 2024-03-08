package main

import (
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

	// 定義したhelloHandlerを使うように登録
	http.HandleFunc("/hello", helloHandler)

	// サーバ起動時のログ出力
	log.Println("server start at port 8080")

	// ListenAndServe関数にて、サーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}