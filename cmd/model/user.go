package model

import "time"

// User /*
/*
ToDo: Implement proper auth
*/
type User struct {
	Id               int       `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Password         string    `json:"password" db:"password"`
	Description      string    `json:"description" db:"description"`
	AvatarPath       string    `json:"avatarPath" db:"avatar_path"`
	Banner           Image     `json:"banner" db:"banner"`
	NumberOfPosts    int       `json:"numberOfPosts" db:"number_of_posts"`
	NumberOfComments int       `json:"numberOfComments" db:"number_of_comments"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt" db:"updated_at"`
}
