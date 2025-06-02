package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/storage"
)

func GetCategories() ([]model.Category, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM categories;`

	var categories []model.Category

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c model.Category
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

func GetCategoryById(id string) (*model.Category, error) {
	db := storage.GetDB()
	var c model.Category
	sqlStatement := `SELECT * FROM categories WHERE id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(&c.Id, &c.Title)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func CreateCategory(c model.Category) (*model.Category, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO categories (title) VALUES ($1) RETURNING id;`
	err := db.QueryRow(sqlStatement, c.Title).Scan(&c.Id)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func UpdateCategory(c model.Category, id string) (*model.Category, error) {
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
