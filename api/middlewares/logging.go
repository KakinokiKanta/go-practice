package middlewares

import (
	"log"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/common"
)

// 自作ResponseWriterを作成
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタ
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter構造体のcodeフィールドに、使うレスポンスコードを保存する
	rsw.code = code

	// HTTPレスポンスに使うレスポンスコードを指定
	// 本来のWriteHeaderメソッドの機能を呼び出し
	rsw.ResponseWriter.WriteHeader(code)
}

// ロギングミドルウェア
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n",traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		// 自作ResponseWriterからロギングしたいデータを出す
		log.Printf("[%d]res: %d\n",traceID, rlw.code)
	})
}
