package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/model"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (*model.Category, error)
	GetAllByCommunityId(id int) ([]model.Category, error)
	Create(c model.Category) (*model.Category, error)
	Update(c model.Category, id int) (*model.Category, error)
	DeleteById(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

var categorySelect = "Categories.id, Categories.title"

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	sqlStatement := `SELECT ` + categorySelect + ` FROM categories;`

	var categories []model.Category

	rows, err := r.db.Query(sqlStatement)

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

func (r *categoryRepository) GetById(id int) (*model.Category, error) {
	var c model.Category
	sqlStatement := `SELECT ` + categorySelect + ` FROM categories WHERE id = $1;`
	err := r.db.QueryRow(sqlStatement, id).Scan(&c.Id, &c.Title)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryRepository) GetAllByCommunityId(id int) ([]model.Category, error) {
	sqlStatement :=
		`SELECT ` + categorySelect + ` FROM Categories
    	 JOIN CommunityCategory ON Categories.id = CommunityCategory.category_id
         WHERE CommunityCategory.community_id = $1;`
	rows, err := r.db.Query(sqlStatement, id)

	if err != nil {
		return nil, err
	}

	var categories []model.Category

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

	return categories, nil
}

func (r *categoryRepository) Create(c model.Category) (*model.Category, error) {
	sqlStatement := `INSERT INTO categories (title) VALUES ($1) RETURNING id;`
	err := r.db.QueryRow(sqlStatement, c.Title).Scan(&c.Id)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryRepository) Update(c model.Category, id int) (*model.Category, error) {
	sqlStatement := `UPDATE categories SET title = $1 WHERE id = $2 RETURNING id;`
	err := r.db.QueryRow(sqlStatement, c.Title, id).Scan(&c.Id)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func (r *categoryRepository) DeleteById(id int) error {
	sqlStatement := `DELETE FROM categories WHERE id = $1;`
	_, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
