package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/KakinokiKanta/go-intermediate/repositories"
	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	// データベース接続のための情報を定義
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// データベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 記事IDが1番の記事内容
	expected := models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum: 2,
	}

	// 記事IDが expected.ID=1の記事を取得して、変数gotに格納
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	// 記事IDが同じかどうか比較
	if got.ID != expected.ID {
		t.Errorf("ID: got %d but want %d\n", got.ID, expected.ID)
	}
	// 記事タイトルが同じかどうか比較
	if got.Title != expected.Title {
		t.Errorf("Title: got %s but want %s\n", got.Title, expected.Title)
	}
	// 記事本文が同じかどうか比較
	if got.Contents != expected.Contents {
		t.Errorf("Contents: got %s but want %s\n", got.Contents, expected.Contents)
	}
	// 記事の投稿者名が同じかどうか比較
	if got.UserName != expected.UserName {
		t.Errorf("UserName: got %s but want %s\n", got.UserName, expected.UserName)
	}
	// 記事いいね数が同じかどうか比較
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
