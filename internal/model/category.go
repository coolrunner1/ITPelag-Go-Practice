package model

import (
	"github.com/coolrunner1/project/internal/storage"
)

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (c *Category) ScanFromRow(row storage.Scannable) error {
	return row.Scan(&c.Id, &c.Title)
}
