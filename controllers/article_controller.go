package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/KakinokiKanta/go-intermediate/apperrors"
	"github.com/KakinokiKanta/go-intermediate/common"
	"github.com/KakinokiKanta/go-intermediate/controllers/services"
	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/gorilla/mux"
)

// Article用のコントローラ構造体
type ArticleController struct {
	service services.ArticleServicer
}

// Article用のコンストラクタ関数
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

/*
POST /article
ブログ記事の投稿をするためのハンドラ
*/
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	authedUserName := common.GetUserName(req.Context())
	if reqArticle.UserName != authedUserName {
		err := apperrors.NotMatchUser.Wrap(errors.New("does not match reqBody user and idtoken user"), "invalid parameter")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

/*
GET /article/list
ブログ記事の一覧を取得するためのハンドラ
*/
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

/*
GET /article/{id}
指定した記事ナンバーの投稿データを取得するためのハンドラ
*/
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// muxを用いてリクエストからidを抽出
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

/*
POST /article/nice
記事にいいねをつけるためのハンドラ
*/
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}
