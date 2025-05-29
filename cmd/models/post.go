package models

import "time"

type Post struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	AuthorId    int        `json:"authorId"`
	CommunityId int        `json:"communityId"`
	Images      []Image    `json:"images"`
	Categories  []Category `json:"categories"`
	Tags        []string   `json:"tags"`
	Views       int        `json:"views"`
	Likes       int        `json:"likes"`
	Dislikes    int        `json:"dislikes"`
	CreatedAt   time.Time  `json:"createdAt"`
}
