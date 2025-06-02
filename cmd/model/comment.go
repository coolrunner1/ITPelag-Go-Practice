package model

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	AuthorId  int       `json:"authorId"`
	PostId    int       `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentInPost struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	AuthorId  int       `json:"authorId"`
	PostId    int       `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
	Author    *User     `json:"author"`
}

type CommentInUserProfile struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	AuthorId  int       `json:"authorId"`
	PostId    int       `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
	Post      *Post     `json:"author"`
}
