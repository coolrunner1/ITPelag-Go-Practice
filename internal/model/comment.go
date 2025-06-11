package model

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	AuthorId  int       `json:"authorId"`
	PostId    int       `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
}
