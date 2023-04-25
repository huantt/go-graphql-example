package resolver

import "github.com/huantt/go-graphql-sample/model"

type UserResolver struct {
	user *model.User
}

func NewUserResolver(user *model.User) *UserResolver {
	return &UserResolver{
		user: user,
	}
}

func (r *UserResolver) Id() *int32 {
	return &r.user.Id
}

func (r *UserResolver) Name() *string {
	return &r.user.Name
}
