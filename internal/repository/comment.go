package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/storage"
)

func GetComments() ([]model.Comment, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM comments;`

	var comments []model.Comment

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c model.Comment
		err := rows.Scan(
			&c.Id,
			&c.Content,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	if len(comments) == 0 {
		return nil, sql.ErrNoRows
	}

	return comments, nil
}
