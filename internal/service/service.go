package service

import (
	"OdinVOdin/internal/mode"
	"OdinVOdin/internal/models"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Posts
	Comments
}

func NewServices(gateways *mode.Modes, logger *logrus.Logger) *Services {
	return &Services{
		Posts:    NewPostsService(gateways.Posts, logger),
		Comments: NewCommentsService(gateways.Comments, logger, gateways.Posts),
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(page, pageSize *int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error)
	GetRepliesOfComment(commentId int) ([]*models.Comment, error)
}
