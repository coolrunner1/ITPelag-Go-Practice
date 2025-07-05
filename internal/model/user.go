package model

import (
	"github.com/LukaGiorgadze/gonull/v2"
	"time"
)

type User struct {
	Id               int                     `json:"id" db:"id"`
	Username         string                  `json:"username" db:"username"`
	Password         string                  `json:"-" db:"password"`
	Email            string                  `json:"email" db:"email"`
	Description      gonull.Nullable[string] `json:"description" db:"description"`
	AvatarPath       gonull.Nullable[string] `json:"avatarPath" db:"avatar_path"`
	BannerPath       gonull.Nullable[string] `json:"bannerPath" db:"banner_path"`
	NumberOfPosts    int                     `json:"numberOfPosts" db:"number_of_posts"`
	NumberOfComments int                     `json:"numberOfComments" db:"number_of_comments"`
	CreatedAt        time.Time               `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time               `json:"updatedAt" db:"updated_at"`
}
