package api

import (
	"database/sql"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/api/middlewares"
	"github.com/KakinokiKanta/go-intermediate/controllers"
	"github.com/KakinokiKanta/go-intermediate/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	// ブログ記事の投稿をするためのエンドポイント
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	// ブログ記事の一覧を取得するためのエンドポイント
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	// 指定した記事ナンバーの投稿データを取得するためのエンドポイント
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	// 記事にいいねをつけるためのエンドポイント
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	// 記事にコメントを投稿するためのエンドポイント
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)

	return r
}
