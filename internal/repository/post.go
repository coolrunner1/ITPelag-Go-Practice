package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
)

// ToDo: add images
type PostRepository interface {
	GetAll(start, limit int) (*dto.PostSearchResponse, error)
	GetById(id int) (*model.Post, error)
	GetAllByUserId(userID int) (*model.Post, error)
	GetAllByCommunityId(start, limit, communityId int) (*dto.PostSearchResponse, error)
	GetAllTagsByPostId(id int) ([]string, error)
	Create(req *dto.PostCreateRequest, userId, communityId int) (*model.Post, error)
	Update(user model.User) (*model.Post, error)
	DeleteById(id int) error
}

type postRepository struct {
	db           *sql.DB
	categoryRepo CategoryRepository
	userRepo     UserRepository
}

func NewPostRepository(db *sql.DB, categoryRepository CategoryRepository, userRepository UserRepository) PostRepository {
	return &postRepository{
		db:           db,
		categoryRepo: categoryRepository,
		userRepo:     userRepository,
	}
}

var postSelect = "id, title, content, author_id, community_id, views, likes, dislikes, created_at, updated_at"

func (r *postRepository) GetAll(start, limit int) (*dto.PostSearchResponse, error) {
	if start < 0 {
		start = 0
	}

	if limit < 0 {
		limit = 15
	}

	sqlStatement := `SELECT ` + postSelect + ` FROM Posts LIMIT $1 OFFSET $2;`

	var posts []model.Post

	rows, err := r.db.Query(sqlStatement, limit, start)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post model.Post
		err := post.ScanFromRow(rows)
		if err != nil {
			return nil, err
		}

		categories, err := r.categoryRepo.GetAllByCommunityId(post.CommunityId)
		if err != nil {
			return nil, err
		}

		post.Categories = categories

		tags, err := r.GetAllTagsByPostId(post.Id)
		if err != nil {
			return nil, err
		}

		post.Tags = tags

		user, err := r.userRepo.GetById(post.AuthorId)
		if err != nil {
			return nil, err
		}

		post.Author = *user

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, sql.ErrNoRows
	}

	var resp dto.PostSearchResponse

	err = r.db.QueryRow("SELECT COUNT(*) FROM Posts").Scan(&resp.Total)

	if err != nil {
		return nil, err
	}

	resp.Data = posts

	return &resp, nil
}

func (r *postRepository) GetById(id int) (*model.Post, error) {
	var post model.Post
	err := post.ScanFromRow(r.db.QueryRow("SELECT "+postSelect+" FROM Posts WHERE id=$1", id))
	if err != nil {
		return nil, err
	}

	categories, err := r.categoryRepo.GetAllByCommunityId(post.CommunityId)
	if err != nil {
		return nil, err
	}

	post.Categories = categories

	tags, err := r.GetAllTagsByPostId(post.Id)
	if err != nil {
		return nil, err
	}

	post.Tags = tags

	user, err := r.userRepo.GetById(post.AuthorId)
	if err != nil {
		return nil, err
	}

	post.Author = *user

	return &post, nil
}

func (r *postRepository) GetAllByUserId(userID int) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) GetAllByCommunityId(start, limit, communityId int) (*dto.PostSearchResponse, error) {
	if start < 0 {
		start = 0
	}

	if limit < 0 {
		limit = 15
	}

	sqlStatement := `SELECT ` + postSelect + ` FROM Posts WHERE community_id = $1 LIMIT $2 OFFSET $3;`

	var posts []model.Post

	rows, err := r.db.Query(sqlStatement, communityId, limit, start)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post model.Post
		err := post.ScanFromRow(rows)
		if err != nil {
			return nil, err
		}

		categories, err := r.categoryRepo.GetAllByCommunityId(post.CommunityId)
		if err != nil {
			return nil, err
		}

		post.Categories = categories

		tags, err := r.GetAllTagsByPostId(post.Id)
		if err != nil {
			return nil, err
		}

		post.Tags = tags

		user, err := r.userRepo.GetById(post.AuthorId)
		if err != nil {
			return nil, err
		}

		post.Author = *user

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, sql.ErrNoRows
	}

	var resp dto.PostSearchResponse

	err = r.db.QueryRow("SELECT COUNT(*) FROM Posts WHERE community_id = $1", communityId).Scan(&resp.Total)

	if err != nil {
		return nil, err
	}

	resp.Data = posts

	return &resp, nil
}

func (r *postRepository) Create(req *dto.PostCreateRequest, userId, communityId int) (*model.Post, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			rollErr := tx.Rollback()
			if rollErr != nil {
				return
			}
		}
	}()

	post := &model.Post{}

	postSQLStatement :=
		`INSERT INTO Posts (title, content, author_id, community_id) 
		 VALUES ($1, $2, $3, $4)
		 RETURNING *;`

	err = post.ScanFromRow(tx.QueryRow(postSQLStatement, req.Title, req.Content, userId, communityId))
	if err != nil {
		return nil, err
	}

	if len(req.Tags) > 0 {
		for _, tag := range req.Tags {
			_, err := tx.Exec("INSERT INTO PostTag (post_id, tag) VALUES ($1, $2)", post.Id, tag)
			if err != nil {
				return nil, err
			}
		}
		post.Tags = req.Tags
	}

	author, err := r.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	post.Author = *author

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) Update(user model.User) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) DeleteById(id int) error {
	_, err := r.db.Exec("DELETE FROM Posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *postRepository) GetAllTagsByPostId(id int) ([]string, error) {
	sqlStatement := `SELECT (tag) FROM PostTag WHERE post_id = $1;`

	rows, err := r.db.Query(sqlStatement, id)

	if err != nil {
		return nil, err
	}

	var tags []string

	for rows.Next() {
		var tag string
		err := rows.Scan(
			&tag,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
