package dto

import "github.com/coolrunner1/project/internal/model"

type UserUpdateRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Description string `json:"description" validate:"omitempty"`
}

type UserSearchResponse struct {
	Total int          `json:"total"`
	Data  []model.User `json:"data"`
}
