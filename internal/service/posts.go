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

type PostsService struct {
	repo   mode.Posts
	logger *logrus.Logger
}

func NewPostsService(repo mode.Posts, logger *logrus.Logger) *PostsService {
	return &PostsService{repo: repo, logger: logger}
}

func (p PostsService) CreatePost(post models.Post) (models.Post, error) {

	if len(post.Author) == 0 {
		p.logger.Println(errors_project.EmptyAuthorError)
		return models.Post{}, errors_project.ResponseError{
			Message: errors_project.EmptyAuthorError,
			Type:    errors_project.BadRequestType,
		}
	}

	if len(post.Content) >= constants.MaxContentLength {
		p.logger.Println(errors_project.TooLongContentError, len(post.Content))
		return models.Post{}, errors_project.ResponseError{
			Message: errors_project.TooLongContentError,
			Type:    errors_project.BadRequestType,
		}
	}

	newPost, err := p.repo.CreatePost(post)
	if err != nil {
		p.logger.Println(errors_project.CreatingPostError, err.Error())
		return models.Post{}, errors_project.ResponseError{
			Message: errors_project.CreatingPostError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return newPost, nil

}

func (p PostsService) GetPostById(postId int) (models.Post, error) {

	if postId <= 0 {
		p.logger.Println(errors_project.WrongIdError, postId)
		return models.Post{}, errors_project.ResponseError{
			Message: errors_project.WrongIdError,
			Type:    errors_project.BadRequestType,
		}
	}

	post, err := p.repo.GetPostById(postId)
	if err != nil {

		p.logger.Println(errors_project.GettingPostError, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return models.Post{}, errors_project.ResponseError{
				Message: errors_project.PostNotFountError,
				Type:    errors_project.NotFoundType,
			}
		}
		return models.Post{}, errors_project.ResponseError{
			Message: errors_project.GettingPostError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return post, nil
}

func (p PostsService) GetAllPosts(page, pageSize *int) ([]models.Post, error) {

	if page != nil && *page < 0 {
		p.logger.Println(errors_project.WrongPageError, *page)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongPageError,
			Type:    errors_project.BadRequestType,
		}
	}

	if pageSize != nil && *pageSize < 0 {
		p.logger.Println(errors_project.WrongPageSizeError, *pageSize)
		return nil, errors_project.ResponseError{
			Message: errors_project.WrongPageSizeError,
			Type:    errors_project.BadRequestType,
		}
	}

	offset, limit := pagination.GetOffsetAndLimit(page, pageSize)

	posts, err := p.repo.GetAllPosts(limit, offset)
	if err != nil {
		p.logger.Println(errors_project.GettingPostError, err.Error())
		return nil, errors_project.ResponseError{
			Message: errors_project.GettingPostError,
			Type:    errors_project.InternalErrorType,
		}
	}

	return posts, nil
}
