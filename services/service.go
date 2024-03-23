package services

import "database/sql"

// MyAppService構造体
type MyAppService struct {
	db *sql.DB
}

// MyAppServiceのコンストラクタ
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
