package graphql

import "OdinVOdin/internal/service"

type Resolver struct {
	PostsService      service.Posts
	CommentsService   service.Comments
	CommentsObservers Observers
}
