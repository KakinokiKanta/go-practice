package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// データベース接続のための情報を定義
var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

// DB接続
func connectDB() (*sql.DB, error) {
	loadEnv()
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadEnv() {
	err := godotenv.Load("./services/.env")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	if err != nil {
		fmt.Println("fail to load .env file")
	}
}
