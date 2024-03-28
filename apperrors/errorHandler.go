package apperrors

import "net/http"

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// エラーの種類を判別して、適切なhttpレスポンスを返す
}
