package dto

// CategoryRequest represent create and update category request body
type CategoryRequest struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}
