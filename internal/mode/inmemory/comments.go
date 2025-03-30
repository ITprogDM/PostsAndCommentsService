package inmemory

import (
	"OdinVOdin/internal/errors_project"
	"OdinVOdin/internal/models"
	"errors"
	"sync"
	"time"
)

type CommentsInMemory struct {
	commCounter int
	Comments    []models.Comment
	mu          sync.RWMutex
}

func NewCommentsInMemory(count int) *CommentsInMemory {
	return &CommentsInMemory{
		commCounter: 0,
		Comments:    make([]models.Comment, 0, count),
	}
}

func (c *CommentsInMemory) CreateComment(comment models.Comment) (models.Comment, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.commCounter++
	t := time.Now()

	comment.ID = c.commCounter
	comment.CreatedAt = t

	c.Comments = append(c.Comments, comment)

	return comment, nil

}

func (c *CommentsInMemory) GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	var res []*models.Comment

	for _, comment := range c.Comments {
		if comment.ReplyTo == nil && comment.Post == postId {
			com := comment
			res = append(res, &com)
		}
	}

	if offset > len(res) {
		return nil, nil
	}

	if offset+limit > len(res) || limit == -1 {
		return res[offset:], nil
	}

	if offset < 0 || limit < 0 {
		return nil, errors.New(errors_project.WrongLimitOffsetError)
	}

	return res[offset : offset+limit], nil
}

func (c *CommentsInMemory) GetRepliesOfComment(commentId int) ([]*models.Comment, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if commentId > c.commCounter {
		return nil, nil
	}

	var res []*models.Comment

	for _, comment := range c.Comments {
		if comment.ReplyTo != nil && *comment.ReplyTo == commentId {
			com := comment
			res = append(res, &com)
		}
	}

	return res, nil
}
