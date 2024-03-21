package services

import (
	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/KakinokiKanta/go-intermediate/repositories"
)

// 指定IDの記事情報を返却
func GetArticleService(articleID int) (models.Article, error) {
	// TODO: sql.DB型を手に入れて、変数dbに代入する

	// repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// repositories層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectedCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 得たコメント一覧commentListを、上記のArticle構造体articleに紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}
