package model

import "time"

type Community struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	BannerPath      string `json:"bannerPath"`
	OwnerID         int    `json:"ownerId"`
	NumberOfMembers int    `json:"numberOfMembers"`
	NumberOfPosts   int    `json:"numberOfPosts"`
	AvatarPath      string `json:"avatarPath"`
	Owner           User   `json:"owner"`
	//Moderators      []User     `json:"moderators"`
	Tags       []string   `json:"tags"`
	Categories []Category `json:"categories"`
	CreatedAt  time.Time  `json:"createdAt"`
}
