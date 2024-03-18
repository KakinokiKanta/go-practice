package repositories_test

import (
	"testing"

	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/KakinokiKanta/go-intermediate/repositories"
	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数を実行
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// SelectArticleList関数から得たArticleスライスの長さが期待通りでないならFAILにする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string // テストのタイトル
		expected models.Article // テストで期待する値
	}{
		{
			// 記事IDが1番の記事内容
			testTitle: "subtest1",
			expected : models.Article{
				ID: 1,
				Title: "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum: 2,
			},
		},
		{
			// 記事IDが2番の記事内容
			testTitle: "subtest2",
			expected : models.Article{
				ID: 2,
				Title: "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum: 4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			// 記事IDが test.expected.IDの記事を取得して、変数gotに格納
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// 記事IDが同じかどうか比較
			if got.ID != test.expected.ID {
				t.Errorf("ID: got %d but want %d\n", got.ID, test.expected.ID)
			}
			// 記事タイトルが同じかどうか比較
			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}
			// 記事本文が同じかどうか比較
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: got %s but want %s\n", got.Contents, test.expected.Contents)
			}
			// 記事の投稿者名が同じかどうか比較
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: got %s but want %s\n", got.UserName, test.expected.UserName)
			}
			// 記事いいね数が同じかどうか比較
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}
