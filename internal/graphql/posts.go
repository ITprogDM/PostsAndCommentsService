package graphql

import (
	"OdinVOdin/graph"
	"OdinVOdin/internal/errors_project"
	"OdinVOdin/internal/models"
	"context"
	"errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreatePost(ctx context.Context, post models.InputPost) (*models.PostGraph, error) {
	newPost, err := r.PostsService.CreatePost(post.FromInput())
	if err != nil {
		var rErr errors_project.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	return newPost.ToGraph(), nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *models.Post, page *int, pageSize *int) ([]*models.Comment, error) {
	comments, err := r.CommentsService.GetCommentsByPost(obj.ID, page, pageSize)
	if err != nil {
		var rErr errors_project.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	return comments, nil
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context, page *int, pageSize *int) ([]*models.PostGraph, error) {
	posts, err := r.PostsService.GetAllPosts(page, pageSize)
	if err != nil {
		var rErr errors_project.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	return models.ToPostGraph(posts), nil
}

// GetPostByID is the resolver for the GetPostById field.
func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	post, err := r.PostsService.GetPostById(id)
	if err != nil {
		var rErr errors_project.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	return &post, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
