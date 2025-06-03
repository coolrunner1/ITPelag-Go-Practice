package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/storage"
)

type UserRepository interface {
	GetAll(start, limit int) ([]model.User, error)
	GetById(id int) (*model.User, error)
	Create(c model.User) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	//Update(c model.Category, id int) (*model.Category, error)
	//DeleteById(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(start, limit int) ([]model.User, error) {
	if start < 0 {
		start = 0
	}

	if limit < 0 {
		limit = 15
	}

	sqlStatement := `SELECT * FROM users LIMIT $1 OFFSET $2;`

	var users []model.User

	rows, err := r.db.Query(sqlStatement, limit, start)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.BannerId,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.Description,
			&user.AvatarPath,
			&user.NumberOfComments,
			&user.NumberOfPosts,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}

	return users, nil
}

func (r *userRepository) GetById(id int) (*model.User, error) {
	db := storage.GetDB()
	var user model.User
	sqlStatement := `SELECT * FROM users WHERE id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(
		&user.Id,
		&user.BannerId,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Description,
		&user.AvatarPath,
		&user.NumberOfComments,
		&user.NumberOfPosts,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user model.User) (*model.User, error) {
	sqlStatement := `INSERT INTO users (username, email, password, description) VALUES ($1, $2, $3, $4) RETURNING id;`
	err := r.db.QueryRow(sqlStatement, user.Username, user.Email, user.Password, user.Description).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	sqlStatement := `SELECT * FROM users WHERE username = $1;`
	var user model.User
	err := r.db.QueryRow(sqlStatement, username).Scan(
		&user.Id,
		&user.BannerId,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Description,
		&user.AvatarPath,
		&user.NumberOfComments,
		&user.NumberOfPosts,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	sqlStatement := `SELECT * FROM users WHERE email = $1;`
	var user model.User
	err := r.db.QueryRow(sqlStatement, email).Scan(
		&user.Id,
		&user.BannerId,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Description,
		&user.AvatarPath,
		&user.NumberOfComments,
		&user.NumberOfPosts,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
