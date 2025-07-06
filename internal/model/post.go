package model

import "time"

type Post struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	AuthorId    int        `json:"authorId"`
	CommunityId int        `json:"communityId"`
	Views       int        `json:"views"`
	Likes       int        `json:"likes"`
	Dislikes    int        `json:"dislikes"`
	Images      []Image    `json:"images"`
	Categories  []Category `json:"categories"`
	Tags        []string   `json:"tags"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
