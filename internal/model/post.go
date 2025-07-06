package model

import (
	"github.com/coolrunner1/project/internal/storage"
	"time"
)

type Post struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	AuthorId    int        `json:"authorId"`
	CommunityId int        `json:"communityId"`
	Views       int        `json:"views"`
	Likes       int        `json:"likes"`
	Dislikes    int        `json:"dislikes"`
	Author      User       `json:"author"`
	Images      []Image    `json:"images"`
	Categories  []Category `json:"categories"`
	Tags        []string   `json:"tags"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (p *Post) ScanFromRow(row storage.Scannable) error {
	return row.Scan(
		&p.Id,
		&p.Title,
		&p.Content,
		&p.AuthorId,
		&p.CommunityId,
		&p.Views,
		&p.Likes,
		&p.Dislikes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}
