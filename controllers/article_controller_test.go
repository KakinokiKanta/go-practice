package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestArticleListHandler(t *testing.T) {
	// テスト対象のハンドラメソッドに入れるinputを定義
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{
			name:       "number query",
			query:      "1",
			resultCode: http.StatusOK,
		},
		{
			name:       "alphabet query",
			query:      "aaa",
			resultCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハンドラに渡す2つの引数 w http.ResponseWriter, req *http.Request を用意する
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			// テスト対象のハンドラメソッドからoutputを得る
			aCon.ArticleListHandler(res, req)

			// outputが期待通りがチェック
			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {
	// テストケースを定義
	var tests = []struct {
		name       string
		articleID  string
		resultCode int
	}{
		{
			name:       "number pathparam",
			articleID:  "1",
			resultCode: http.StatusOK,
		},
		{
			name:       "alphabet pathparam",
			articleID:  "aaa",
			resultCode: http.StatusNotFound,
		},
	}

	// テーブルドリブンに実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest.NewRequest 関数でリクエストを作成
			url := fmt.Sprintf("http://localhost:8080/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			// httptest.ResponseRecorder構造体を用意
			res := httptest.NewRecorder()

			// ハンドラメソッドの実行
			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
