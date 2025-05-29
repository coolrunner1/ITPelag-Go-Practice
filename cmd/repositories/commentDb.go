package repositories

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/models"
	"github.com/coolrunner1/project/cmd/storage"
)

func GetComments() ([]models.Comment, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM comments;`

	var comments []models.Comment

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Comment
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
