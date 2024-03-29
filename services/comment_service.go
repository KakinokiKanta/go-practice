package services

import (
	"github.com/KakinokiKanta/go-intermediate/apperrors"
	"github.com/KakinokiKanta/go-intermediate/models"
	"github.com/KakinokiKanta/go-intermediate/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
