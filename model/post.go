package model

type Post struct {
	Id       int32     `json:"id"`
	UserId   int32     `json:"userId"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comments []Comment `json:"comments"`
}
