package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// テスト全体で共有するsql.DB型
var testDB *sql.DB

// データベース接続のための情報を定義
var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

// 前処理
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("setup")
		return err
	}
	if err := setupTestData(); err != nil {
		return err
	}

	return nil
}

// DB接続
func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return nil
}

// データベース内のテスト用データを消す後処理
func cleanupDB() error {
	// os/execパッケージのexec.Command関数を用いて、実行したいコマンドの情報を持つexec.Cmd型の変数を用意
	cmd := exec.Command("mysql", "-h", "localhost", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	// exec.Cmd型のRunメソッドを読んで、コマンドを実行
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// データベースにデータを入れる前処理
func setupTestData() error {
	// os/execパッケージのexec.Command関数を用いて、実行したいコマンドの情報を持つexec.Cmd型の変数を用意
	cmd := exec.Command("mysql", "-h", "localhost", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setupDB.sql")
	// exec.Cmd型のRunメソッドを読んで、コマンドを実行
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// 接続したデータベースとのアクセスを閉じる後処理
func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup() // 前処理
	if err != nil {
		os.Exit(1)
	}

	m.Run() // パッケージ内のユニットテストすべてを実行

	teardown() // 後処理
}
