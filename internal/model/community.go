package model

import (
	"github.com/LukaGiorgadze/gonull/v2"
	"github.com/coolrunner1/project/internal/storage"
	"time"
)

type Community struct {
	ID              int                     `json:"id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	BannerPath      gonull.Nullable[string] `json:"bannerPath"`
	AvatarPath      gonull.Nullable[string] `json:"avatarPath"`
	OwnerID         int                     `json:"ownerId"`
	NumberOfMembers int                     `json:"numberOfMembers"`
	NumberOfPosts   int                     `json:"numberOfPosts"`
	Owner           User                    `json:"owner"`
	//Moderators      []User     `json:"moderators"`
	Tags       []string   `json:"tags"`
	Categories []Category `json:"categories"`
	CreatedAt  time.Time  `json:"createdAt"`
}

func (community *Community) ScanFromRow(row storage.Scannable) error {
	return row.Scan(
		&community.ID,
		&community.Name,
		&community.Description,
		&community.BannerPath,
		&community.AvatarPath,
		&community.OwnerID,
		&community.NumberOfMembers,
		&community.NumberOfPosts,
		&community.CreatedAt,
	)
}
