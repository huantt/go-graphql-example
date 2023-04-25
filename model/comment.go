package model

type Comment struct {
	Id     int32  `json:"id"`
	PostId int32  `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
