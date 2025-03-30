package mode

import "OdinVOdin/internal/models"

type Modes struct {
	Posts
	Comments
}

func NewModes(posts Posts, comments Comments) *Modes {
	return &Modes{
		Posts:    posts,
		Comments: comments,
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(limit, offset int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error)
	GetRepliesOfComment(commentId int) ([]*models.Comment, error)
}
