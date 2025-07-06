package dto

import (
	"github.com/coolrunner1/project/internal/model"
)

type CommunityCreateRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Description string   `json:"description" validate:"min=3,max=500"`
	Tags        []string `json:"tags"`
	Categories  []int    `json:"categories"`
}

type CommunityUpdateRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Description string   `json:"description" validate:"min=3,max=500"`
	OwnerID     int      `json:"ownerId"`
	Tags        []string `json:"tags"`
	Categories  []int    `json:"categories"`
}

type CommunitySearchResponse struct {
	Total int               `json:"total"`
	Data  []model.Community `json:"data"`
}
