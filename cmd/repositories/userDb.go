package repositories

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/models"
	"github.com/coolrunner1/project/cmd/storage"
)

func GetUsers() ([]models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM users;`

	var users []models.User

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
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

func GetUserById(id int) (*models.User, error) {
	db := storage.GetDB()
	var user models.User
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
func CreateUser(user models.User) (*models.User, error) {
	/*db := storage.GetDB()
	sqlStatement := `INSERT INTO users (username, password, description, createdAt) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id;`
	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.Description).Scan(&user.Id)

	if err != nil {
		return nil, err
	}
	*/
	return &user, nil
}
