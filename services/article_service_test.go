package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/KakinokiKanta/go-intermediate/services"
	_ "github.com/go-sql-driver/mysql"
)

// レシーバとして使うためのMyAppService構造体
var aSer *services.MyAppService

func TestMain(m *testing.M) {
	// sql.DB型を作成
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sql.DB型からサービス構造体を作成
	aSer = services.NewMyAppService(db)

	// 個別のベンチマークテストの実行
	m.Run()
}

// GetArticleServiceメソッドのベンチマーク用のテスト
func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}