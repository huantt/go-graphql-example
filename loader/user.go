package loader

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/huantt/go-graphql-sample/model"
	"strconv"
)

type UserLoader struct {
	userRepository UserRepository
}

func NewUserLoader(userRepository UserRepository) *UserLoader {
	return &UserLoader{
		userRepository: userRepository,
	}
}

type UserRepository interface {
	FindUsers(ctx context.Context, ids []int32) ([]model.User, error)
}

func (c *UserLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var ids []int32
	for i := range keys {
		id, err := strconv.Atoi(keys[i].String())
		if err != nil {
			loadBatchError(err, i)
		}
		ids = append(ids, int32(id))
	}

	var records = make([]*dataloader.Result, len(keys))
	users, err := c.userRepository.FindUsers(ctx, ids)
	if err != nil {
		return loadBatchError(err, len(keys))
	}
	for _, user := range users {
		// Keep order
		i := mustIndex(ids, user.Id)
		records[i] = &dataloader.Result{Data: &users[i]}
	}
	return records
}

func GetUser(ctx context.Context, id int32) (*model.User, error) {
	user, err := loadOne(ctx, userLoaderKey, key(id))
	if err != nil || user == nil {
		return nil, err
	}
	res, ok := user.(*model.User)
	if !ok {
		return nil, fmt.Errorf("wrong type: %T", user)
	}
	return res, nil
}
