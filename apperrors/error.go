package apperrors

// 独自エラーに含めるフィールドの定義
type MyAppError struct {
	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  `json:"-"` // エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
