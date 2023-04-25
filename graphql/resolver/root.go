package resolver

import (
	"context"
	"github.com/huantt/go-graphql-sample/model"
)

// The RootQueryResolver is the entry point for all top-level read operations.
type RootQueryResolver struct {
	postRepository PostRepository
}

func NewRootResolver(postRepository PostRepository) (*RootQueryResolver, error) {
	return &RootQueryResolver{
		postRepository: postRepository,
	}, nil
}

type PostRepository interface {
	GetPosts(ctx context.Context, limit int32) ([]model.Post, error)
	FindPosts(ctx context.Context, ids []int32) ([]model.Post, error)
}

type listPostsArgs struct {
	Limit int32
}

func (r *RootQueryResolver) ListPosts(ctx context.Context, args listPostsArgs) (*[]*PostResolver, error) {
	posts, err := r.postRepository.GetPosts(ctx, args.Limit)
	if err != nil {
		return nil, err
	}
	postResolvers := make([]*PostResolver, len(posts))
	for i := range posts {
		postResolvers[i] = NewPostResolver(&posts[i])
	}
	return &postResolvers, nil
}

type findPostArgs struct {
	Ids []int32
}

func (r *RootQueryResolver) FindPosts(ctx context.Context, args findPostArgs) (*[]*PostResolver, error) {
	posts, err := r.postRepository.FindPosts(ctx, args.Ids)
	if err != nil {
		return nil, err
	}
	postResolvers := make([]*PostResolver, len(posts))
	for i := range posts {
		postResolvers[i] = NewPostResolver(&posts[i])
	}
	return &postResolvers, nil
}
