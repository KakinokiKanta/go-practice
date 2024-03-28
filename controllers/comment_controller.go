package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KakinokiKanta/go-intermediate/apperrors"
	"github.com/KakinokiKanta/go-intermediate/controllers/services"
	"github.com/KakinokiKanta/go-intermediate/models"
)

// Comment用のコントローラ構造体
type CommentController struct {
	service services.CommentServicer
}

// Comment用のコンストラクタ関数
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

/*
POST /comment
記事にコメントを投稿するためのハンドラ
*/
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデータを構造体にデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
