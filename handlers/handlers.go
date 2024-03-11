package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/gorilla/mux"
)

/*
	POST /article
	ブログ記事の投稿をするためのハンドラ
*/
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// デコードした構造体をjsonエンコードしてレスポンスとする
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}


/*
	GET /article/list
	ブログ記事の一覧を取得するためのハンドラ
*/
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// モックデータを呼び出してjsonエンコード
	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	// jsonをレスポンスに格納
	w.Write(jsonData)
}


/*
	GET /article/{id}
	指定した記事ナンバーの投稿データを取得するためのハンドラ
*/
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// muxを用いてリクエストからidを抽出
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	
	// モックデータを呼び出してjsonエンコード
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	// jsonをレスポンスに格納
	w.Write(jsonData)
}


/*
	POST /article/nice
	記事にいいねをつけるためのハンドラ
*/
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// モックデータを呼び出してjsonエンコード
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	// jsonをレスポンスに格納
	w.Write(jsonData)
}


/*
	POST /comment
	記事にコメントを投稿するためのハンドラ
*/
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// モックデータを呼び出してjsonエンコード
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}

	// jsonをレスポンスに格納
	w.Write(jsonData)
}
