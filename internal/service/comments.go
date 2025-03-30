package service

import (
	"OdinVOdin/internal/constants"
	"OdinVOdin/internal/errors_project"
	"OdinVOdin/internal/mode"
	"OdinVOdin/internal/models"
	"OdinVOdin/internal/pagination"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
)

type CommentsService struct {
	repo       mode.Comments
	logger     *logrus.Logger
	PostGetter PostGetter
}

type PostGetter interface {
	GetPostById(id int) (models.Post, error)
}

func NewCommentsService(repo mode.Comments, logger *logrus.Logger, getter PostGetter) *CommentsService {
	return &CommentsService{repo: repo, logger: logger, PostGetter: getter}
}

func (c CommentsService) CreateComment(comment models.Comment) (models.Comment, error) {
	if len(comment.Author) == 0 {
		c.logger.Println(errors_project.EmptyAuthorError)
		return models.Comment{}, errors_project.ResponseError{
			Message: errors_project.EmptyAuthorError,
			Type:    errors_project.BadRequestType,
		}
	}

	if len(comment.Content) >= constants.MaxContentLength {
		c.logger.Println(errors_project.TooLongContentError, len(comment.Content))
		return models.Comment{}, errors_project.ResponseError{
			Message: errors_project.TooLongContentError,
			Type:    errors_project.BadRequestType,
		}
	}

	if comment.Post <= 0 {
		c.logger.Println(errors_project.WrongIdError, comment.Post)
		return models.Comment{}, errors_project.ResponseError{
			Message: errors_project.WrongIdError,
			Type:    errors_project.BadRequestType,
		}
	}

	post, err := c.PostGetter.GetPostById(comment.Post)
	if err != nil {
		c.logger.Println(errors_project.GettingPostError, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return models.Comment{}, errors_project.ResponseError{
				Message: errors_project.PostNotFountError,
				Type:    errors_project.NotFoundType,
			}
		}
	}

	if !post.CommentsAllowed {
		c.logger.Println(errors_project.CommentsNotAllowedError)
		return models.Comment{}, errors_project.ResponseError{
			Message: errors_project.CommentsNotAllowedError,
			Type:    errors_project.BadRequestType,
		}

	}

	newComment, err := c.repo.CreateComment(comment)
	if err != nil {
		c.logger.Println(errors_project.CreatingCommentError, err.Error())
		return models.Comment{}, errors_project.ResponseError{
			Message: errors_project.CreatingCommentError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {

	if postId <= 0 {
		c.logger.Println(errors_project.WrongIdError, postId)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongIdError,
			Type:    errors_project.BadRequestType,
		}
	}

	if page != nil && *page < 0 {
		c.logger.Println(errors_project.WrongPageError, *page)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongPageError,
			Type:    errors_project.BadRequestType,
		}
	}

	if pageSize != nil && *pageSize < 0 {
		c.logger.Println(errors_project.WrongPageSizeError, *pageSize)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongPageSizeError,
			Type:    errors_project.BadRequestType,
		}
	}

	offset, limit := pagination.GetOffsetAndLimit(page, pageSize)

	comments, err := c.repo.GetCommentsByPost(postId, limit, offset)
	if err != nil {
		c.logger.Println(errors_project.GettingCommentError, postId, err.Error())
		return nil, errors_project.ResponseError{
			Message: errors_project.GettingCommentError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return comments, nil
}

func (c CommentsService) GetRepliesOfComment(commentId int) ([]*models.Comment, error) {

	if commentId <= 0 {
		c.logger.Println(errors_project.WrongIdError, commentId)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongIdError,
			Type:    errors_project.BadRequestType,
		}
	}

	comments, err := c.repo.GetRepliesOfComment(commentId)
	if err != nil {
		c.logger.Println(errors_project.GettingRepliesError, commentId, err.Error())
		return nil, errors_project.ResponseError{
			Message: errors_project.GettingRepliesError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return comments, nil

}
