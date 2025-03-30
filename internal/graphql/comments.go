package graphql

import (
	"OdinVOdin/graph"
	"OdinVOdin/internal/errors_project"
	"OdinVOdin/internal/models"
	"context"
	"errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment) ([]*models.Comment, error) {
	comments, err := r.CommentsService.GetRepliesOfComment(obj.ID)
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

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input models.InputComment) (*models.Comment, error) {
	newComment, err := r.CommentsService.CreateComment(input.FromInput())
	if err != nil {
		var rErr errors_project.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	if err := r.CommentsObservers.NotifyObservers(newComment.Post, newComment); err != nil {
		if err.Error() != errors_project.ThereIsNoObserversError {
			return nil, &gqlerror.Error{
				Extensions: map[string]interface{}{
					"message": err,
					"type":    errors_project.InternalErrorType,
				},
			}
		}
	}

	return &newComment, nil
}

// CommentsSubscription is the resolver for the CommentsSubscription field.
func (r *subscriptionResolver) CommentsSubscription(ctx context.Context, postID int) (<-chan *models.Comment, error) {
	id, ch, err := r.CommentsObservers.CreateObserver(postID)

	if err != nil {
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"message": err,
				"type":    errors_project.InternalErrorType,
			},
		}
	}

	go func() {
		<-ctx.Done()
		err := r.CommentsObservers.DeleteObserver(id, postID)
		if err != nil {
			return
		}
	}()

	return ch, nil
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
