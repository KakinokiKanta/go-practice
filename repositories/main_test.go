package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// テスト全体で共有するsql.DB型
var testDB * sql.DB

// 全テスト共通の前処理を書く
func setup() error {
	// データベース接続のための情報を定義
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// データベースに接続
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return nil
}

// 全テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}

func TestMain (m *testing.M) {
	err := setup() // 前処理
	if err != nil {
		os.Exit(1)
	}

	m.Run() // パッケージ内のユニットテストすべてを実行

	teardown() // 後処理
}
