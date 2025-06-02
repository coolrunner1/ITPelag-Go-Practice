package model

import "time"

// User /*
/*
ToDo: Implement proper auth
*/
type User struct {
	Id               int       `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Description      string    `json:"description"`
	AvatarPath       string    `json:"avatarPath"`
	Banner           Image     `json:"banner"`
	NumberOfPosts    int       `json:"numberOfPosts"`
	NumberOfComments int       `json:"numberOfComments"`
	CreatedAt        time.Time `json:"createdAt"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
