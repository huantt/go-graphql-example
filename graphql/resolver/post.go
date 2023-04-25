package resolver

import (
	"context"
	"github.com/huantt/go-graphql-sample/loader"
	"github.com/huantt/go-graphql-sample/model"
)

type PostResolver struct {
	post *model.Post
}

func NewPostResolver(post *model.Post) *PostResolver {
	return &PostResolver{
		post: post,
	}
}

func (r *PostResolver) Id() *int32 {
	return &r.post.Id
}
func (r *PostResolver) UserId() *int32 {
	return &r.post.UserId
}

func (r *PostResolver) Title() *string {
	return &r.post.Title
}

func (r *PostResolver) Body() *string {
	return &r.post.Body
}

func (r *PostResolver) Comments() *[]*CommentResolver {
	return nil
}

func (r *PostResolver) User(ctx context.Context) (*UserResolver, error) {
	user, err := loader.GetUser(ctx, r.post.UserId)
	if err != nil {
		return nil, err
	}
	return NewUserResolver(user), nil
}
