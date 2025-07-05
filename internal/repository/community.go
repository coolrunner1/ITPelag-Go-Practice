package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/model"
)

type CommunityRepository interface {
	GetAll(start, limit int) ([]model.Community, error)
	//GetCommunity(id int) (*model.Community, error)
	GetAllTagsByCommunityId(id int) ([]string, error)
	//CreateCommunity(*model.Community) (*model.Community, error)
	//UpdateCommunity(*model.Community) (*model.Community, error)
	//DeleteCommunity(*model.Community) error
}

type communityRepository struct {
	db           *sql.DB
	categoryRepo CategoryRepository
	userRepo     UserRepository
}

func NewCommunityRepository(db *sql.DB, categoryRepository CategoryRepository, userRepository UserRepository) CommunityRepository {
	return &communityRepository{
		db:           db,
		categoryRepo: categoryRepository,
		userRepo:     userRepository,
	}
}

func (r *communityRepository) GetAll(start, limit int) ([]model.Community, error) {
	if start < 0 {
		start = 0
	}

	if limit < 0 {
		limit = 15
	}

	sqlStatement := `SELECT * FROM Communities LIMIT $1 OFFSET $2;`

	var communities []model.Community

	rows, err := r.db.Query(sqlStatement, limit, start)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var community model.Community
		err := rows.Scan(
			&community.ID,
			&community.Name,
			&community.Description,
			&community.BannerPath,
			&community.AvatarPath,
			&community.OwnerID,
			&community.NumberOfMembers,
			&community.NumberOfPosts,
			&community.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories, err := r.categoryRepo.GetAllByCommunityId(community.ID)

		if err != nil {
			return nil, err
		}

		community.Categories = categories

		tags, err := r.GetAllTagsByCommunityId(community.ID)

		if err != nil {
			return nil, err
		}

		community.Tags = tags

		owner, err := r.userRepo.GetById(community.OwnerID)

		if err != nil {
			return nil, err
		}

		community.Owner = *owner

		communities = append(communities, community)
	}

	if len(communities) == 0 {
		return nil, sql.ErrNoRows
	}

	return communities, nil
}

func (r *communityRepository) GetAllTagsByCommunityId(id int) ([]string, error) {
	sqlStatement := `SELECT (tag) FROM CommunityTag WHERE community_id = $1;`

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
