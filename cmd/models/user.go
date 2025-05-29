package models

import "time"

type User struct {
	Id               int       `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Description      string    `json:"description"`
	AvatarPath       string    `json:"avatarPath"`
	NumberOfPosts    int       `json:"numberOfPosts"`
	NumberOfComments int       `json:"numberOfComments"`
	CreatedAt        time.Time `json:"createdAt"`
}
