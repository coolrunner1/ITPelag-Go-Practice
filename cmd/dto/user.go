package dto

type UserUpdateRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Description string `json:"description" validate:"omitempty"`
}
