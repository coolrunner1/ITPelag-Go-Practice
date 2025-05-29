package models

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title" validate:"required"`
}
