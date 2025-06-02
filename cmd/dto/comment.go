package dto

import (
	"github.com/coolrunner1/project/cmd/model"
	"time"
)

type PostCommentDTO struct {
	Id        int         `json:"id"`
	Content   string      `json:"content"`
	AuthorId  int         `json:"authorId"`
	PostId    int         `json:"postId"`
	CreatedAt time.Time   `json:"createdAt"`
	Author    *model.User `json:"author"`
}

type UserProfileCommentDTO struct {
	Id        int         `json:"id"`
	Content   string      `json:"content"`
	AuthorId  int         `json:"authorId"`
	PostId    int         `json:"postId"`
	CreatedAt time.Time   `json:"createdAt"`
	Post      *model.Post `json:"post"`
}
