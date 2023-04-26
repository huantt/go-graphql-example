package adapter

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/huantt/go-graphql-sample/model"
	"github.com/huantt/go-graphql-sample/pkg/log"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type JsonPlaceHolder struct {
	httpClient *resty.Client
}

func (j *JsonPlaceHolder) FindPosts(ctx context.Context, ids []int32) ([]model.Post, error) {
	log.Debugf("FindPosts: %v", ids)
	var posts []model.Post
	request := j.httpClient.R().SetResult(&posts)
	stringIds := make([]string, 0, len(ids))
	for _, id := range ids {
		stringIds = append(stringIds, fmt.Sprintf("%d", id))
	}
	request.SetQueryParamsFromValues(map[string][]string{
		"id": stringIds,
	})
	resp, err := request.Get("/posts")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}
	return posts, nil
}

const endpoint = "https://jsonplaceholder.typicode.com"

func NewJsonPlaceHolder() *JsonPlaceHolder {
	httpClient := resty.New()
	httpClient.
		SetRetryCount(12).
		SetRetryWaitTime(5 * time.Second).
		SetBaseURL(endpoint).AddRetryCondition(func(response *resty.Response, err error) bool {
		if err != nil {
			return true
		}
		if response.StatusCode() == http.StatusInternalServerError ||
			response.StatusCode() == http.StatusBadGateway ||
			response.StatusCode() == http.StatusGatewayTimeout ||
			response.StatusCode() == http.StatusServiceUnavailable {
			log.Warnf("Response status code is %d - Request: %s - Body: %s - Retrying...", response.StatusCode(), response.Request.URL, response.Body())
			return true
		}

		return false
	})
	return &JsonPlaceHolder{httpClient}
}

func (j *JsonPlaceHolder) GetPosts(ctx context.Context, limit int32) ([]model.Post, error) {
	log.Debugf("GetPosts: %d", limit)
	var posts []model.Post
	resp, err := j.httpClient.R().
		SetQueryParam("limit", fmt.Sprintf("%d", limit)).
		SetResult(&posts).
		Get("/posts")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}
	return posts[0:limit], nil
}

func (j *JsonPlaceHolder) FindUsers(ctx context.Context, ids []int32) ([]model.User, error) {
	log.Debugf("FindUsers: %v", ids)
	var users []model.User
	request := j.httpClient.R().SetResult(&users)
	stringIds := make([]string, 0, len(ids))
	for _, id := range ids {
		stringIds = append(stringIds, fmt.Sprintf("%d", id))
	}
	request.SetQueryParamsFromValues(map[string][]string{
		"id": stringIds,
	})
	resp, err := request.Get("/users")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}
	return users, nil
}
