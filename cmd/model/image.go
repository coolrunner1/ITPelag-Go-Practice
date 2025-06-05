package model

import "time"

type Image struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Alt         string    `json:"alt"`
	Description string    `json:"description"`
	ImagePath   string    `json:"imagePath"`
	CreatedAt   time.Time `json:"createdAt"`
}
