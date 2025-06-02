package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/storage"
)

func GetUsers() ([]model.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM users;`

	var users []model.User

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Description,
			//&user.AvatarPath,
			&user.CreatedAt,
		)
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

func GetUserById(id int) (*model.User, error) {
	db := storage.GetDB()
	var user model.User
	sqlStatement := `SELECT * FROM users WHERE id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(
		&user.Id,
		&user.Username,
		&user.Description,
		//&user.AvatarPath,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser /*
/*
ToDo: Hash password with bcrypt
*/
func CreateUser(user model.User) (*model.User, error) {
	/*db := storage.GetDB()
	sqlStatement := `INSERT INTO users (username, password, description, createdAt) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id;`
	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.Description).Scan(&user.Id)

	if err != nil {
		return nil, err
	}
	*/
	return &user, nil
}
