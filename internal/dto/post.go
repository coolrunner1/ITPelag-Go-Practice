package dto

import (
	"github.com/coolrunner1/project/internal/model"
)

type PostCreateRequest struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required,min=3"`
	Tags    []string `json:"tags"`
}

type PostSearchResponse struct {
	Total int          `json:"total"`
	Data  []model.Post `json:"data"`
}
