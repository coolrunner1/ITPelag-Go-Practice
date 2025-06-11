package dto

import (
	model2 "github.com/coolrunner1/project/internal/model"
	"time"
)

type PostCommentResponse struct {
	Id        int          `json:"id"`
	Content   string       `json:"content"`
	AuthorId  int          `json:"authorId"`
	PostId    int          `json:"postId"`
	CreatedAt time.Time    `json:"createdAt"`
	Author    *model2.User `json:"author"`
}

type UserProfileCommentResponse struct {
	Id        int          `json:"id"`
	Content   string       `json:"content"`
	AuthorId  int          `json:"authorId"`
	PostId    int          `json:"postId"`
	CreatedAt time.Time    `json:"createdAt"`
	Post      *model2.Post `json:"post"`
}
