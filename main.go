package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KakinokiKanta/go-intermediate/controllers"
	"github.com/KakinokiKanta/go-intermediate/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// データベース接続のための情報を定義
var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	ser := services.NewMyAppService(db)
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

// DB接続
func connectDB() (*sql.DB, error) {
	loadEnv()
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadEnv() {
	err := godotenv.Load(".env")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	if err != nil {
		fmt.Println("fail to load .env file")
	}
}
