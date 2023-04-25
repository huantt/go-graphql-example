package resolver

import "github.com/huantt/go-graphql-sample/model"

type CommentResolver struct {
	comment *model.Comment
}

func NewCommentResolver(comment *model.Comment) *CommentResolver {
	return &CommentResolver{
		comment: comment,
	}
}

func (r *CommentResolver) PostId() *int32 {
	return &r.comment.PostId
}

func (r *CommentResolver) Id() *int32 {
	return &r.comment.Id
}

func (r *CommentResolver) Name() *string {
	return &r.comment.Name
}

func (r *CommentResolver) Email() *string {
	return &r.comment.Email
}

func (r *CommentResolver) Body() *string {
	return &r.comment.Body
}
