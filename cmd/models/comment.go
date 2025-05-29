package models

import "time"

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	//AuthorId  int       `json:"author_id"`
	//PostId    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
