package resolver

import (
	graphql "github.com/huantt/go-graphql-sample/graphql/scalar"
	"github.com/huantt/go-graphql-sample/model"
)

type UserResolver struct {
	user *model.User
}

func NewUserResolver(user *model.User) *UserResolver {
	return &UserResolver{
		user: user,
	}
}

func (r *UserResolver) Id() (*graphql.Int64, error) {
	return graphql.NewInt64(r.user.Id)
}

func (r *UserResolver) Name() *string {
	return &r.user.Name
}

func (r *UserResolver) Address() (*graphql.Json, error) {
	return graphql.NewJson(r.user.Address)
}
