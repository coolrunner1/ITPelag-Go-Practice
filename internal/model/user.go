package model

import (
	"github.com/LukaGiorgadze/gonull/v2"
	"github.com/coolrunner1/project/internal/storage"
	"time"
)

type User struct {
	Id               int                        `json:"id"`
	Username         string                     `json:"username"`
	Password         string                     `json:"-"`
	Email            string                     `json:"email"`
	Description      gonull.Nullable[string]    `json:"description"`
	AvatarPath       gonull.Nullable[string]    `json:"avatarPath"`
	BannerPath       gonull.Nullable[string]    `json:"bannerPath"`
	NumberOfPosts    int                        `json:"numberOfPosts"`
	NumberOfComments int                        `json:"numberOfComments"`
	CreatedAt        time.Time                  `json:"createdAt"`
	UpdatedAt        time.Time                  `json:"updatedAt"`
	DeletedAt        gonull.Nullable[time.Time] `json:"deletedAt"`
}

func (u *User) ScanFromRow(row storage.Scannable) error {
	return row.Scan(
		&u.Id,
		&u.Email,
		&u.Username,
		&u.Password,
		&u.Description,
		&u.AvatarPath,
		&u.BannerPath,
		&u.NumberOfComments,
		&u.NumberOfPosts,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
}
