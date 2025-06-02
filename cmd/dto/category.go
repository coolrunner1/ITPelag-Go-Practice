package dto

type CategoryDTO struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}
