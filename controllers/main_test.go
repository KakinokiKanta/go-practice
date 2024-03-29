package controllers_test

import (
	"testing"

	"github.com/KakinokiKanta/go-intermediate/controllers"
	"github.com/KakinokiKanta/go-intermediate/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

// テストに使うリソース (コントローラ構造体) を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
