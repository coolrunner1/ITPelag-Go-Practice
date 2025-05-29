package repositories

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/models"
	"github.com/coolrunner1/project/cmd/storage"
)

func GetCategories() ([]models.Category, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM categories;`

	var categories []models.Category

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&c.Id,
			&c.Title,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if len(categories) == 0 {
		return nil, sql.ErrNoRows
	}

	return categories, nil
}

func GetCategoryById(id string) (*models.Category, error) {
	db := storage.GetDB()
	var c models.Category
	sqlStatement := `SELECT * FROM categories WHERE id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(&c.Id, &c.Title)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func CreateCategory(c models.Category) (*models.Category, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO categories (title) VALUES ($1) RETURNING id;`
	err := db.QueryRow(sqlStatement, c.Title).Scan(&c.Id)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func UpdateCategory(c models.Category, id string) (*models.Category, error) {
	db := storage.GetDB()
	sqlStatement := `UPDATE categories SET title = $1 WHERE id = $2 RETURNING id;`
	err := db.QueryRow(sqlStatement, c.Title, id).Scan(&c.Id)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func DeleteCategoryById(id string) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM categories WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
