package repository

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
)

type CommunityRepository interface {
	GetAll(start, limit int) (*dto.CommunitySearchResponse, error)
	GetById(id int) (*model.Community, error)
	GetAllTagsByCommunityId(id int) ([]string, error)
	Create(req *dto.CommunityCreateRequest, userId int) (*model.Community, error)
	Update(req *dto.CommunityUpdateRequest, userId, communityId int) (*model.Community, error)
	DeleteById(id int) error
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

func (r *communityRepository) GetAll(start, limit int) (*dto.CommunitySearchResponse, error) {
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
		err := community.ScanFromRow(rows)

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

	var resp dto.CommunitySearchResponse

	err = r.db.QueryRow("SELECT COUNT(*) FROM Communities").Scan(&resp.Total)

	if err != nil {
		return nil, err
	}

	resp.Data = communities

	return &resp, nil
}

func (r *communityRepository) GetById(id int) (*model.Community, error) {
	sqlStatement := `SELECT * FROM Communities WHERE id = $1;`

	var community model.Community

	err := community.ScanFromRow(r.db.QueryRow(sqlStatement, id))

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

	return &community, nil
}

func (r *communityRepository) Create(req *dto.CommunityCreateRequest, userId int) (*model.Community, error) {
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

	community := &model.Community{}

	communitySQLStatement :=
		`INSERT INTO Communities (name, description, owner_id) 
		 VALUES ($1, $2, $3)
		 RETURNING *;`

	err = community.ScanFromRow(tx.QueryRow(communitySQLStatement, req.Name, req.Description, userId))
	if err != nil {
		return nil, err
	}

	if len(req.Tags) > 0 {
		for _, tag := range req.Tags {
			_, err := tx.Exec("INSERT INTO CommunityTag (community_id, tag) VALUES ($1, $2)", community.ID, tag)
			if err != nil {
				return nil, err
			}
		}
		community.Tags = req.Tags
	}

	var categories []model.Category

	if len(req.Categories) > 0 {
		for _, categoryID := range req.Categories {
			var category *model.Category
			category, err = r.categoryRepo.GetById(categoryID)
			if err != nil {
				return nil, err
			}
			_, err = tx.Exec(
				`INSERT INTO CommunityCategory (community_id, category_id)
				VALUES ($1, $2)`,
				community.ID, categoryID,
			)
			if err != nil {
				return nil, err
			}
			categories = append(categories, *category)
		}
		community.Categories = categories
	}

	owner, err := r.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	community.Owner = *owner

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return community, nil
}

func (r *communityRepository) Update(req *dto.CommunityUpdateRequest, userId, communityId int) (*model.Community, error) {
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

	if req.OwnerID < 1 {
		req.OwnerID = userId
	}

	community := &model.Community{}

	communitySQLStatement :=
		`UPDATE Communities
		 SET
		     name = $1,
		     description = $2,
		     owner_id = $3
		 WHERE id = $4
		 RETURNING *;`

	err = community.ScanFromRow(tx.QueryRow(communitySQLStatement, req.Name, req.Description, req.OwnerID, communityId))
	if err != nil {
		return nil, err
	}

	if len(req.Tags) > 0 {
		_, err := tx.Exec("DELETE FROM CommunityTag WHERE community_id = $1", communityId)
		if err != nil {
			return nil, err
		}
		for _, tag := range req.Tags {
			_, err := tx.Exec("INSERT INTO CommunityTag (community_id, tag) VALUES ($1, $2)", communityId, tag)
			if err != nil {
				return nil, err
			}
		}
		community.Tags = req.Tags
	}

	var categories []model.Category

	if len(req.Categories) > 0 {
		_, err := tx.Exec("DELETE FROM CommunityCategory WHERE community_id = $1", communityId)
		if err != nil {
			return nil, err
		}
		for _, categoryID := range req.Categories {
			var category *model.Category
			category, err = r.categoryRepo.GetById(categoryID)
			if err != nil {
				return nil, err
			}
			_, err = tx.Exec(
				`INSERT INTO CommunityCategory (community_id, category_id)
				VALUES ($1, $2)`,
				community.ID, categoryID,
			)
			if err != nil {
				return nil, err
			}
			categories = append(categories, *category)
		}
		community.Categories = categories
	}

	owner, err := r.userRepo.GetById(community.OwnerID)
	if err != nil {
		return nil, err
	}

	community.Owner = *owner

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return community, nil
}

func (r *communityRepository) DeleteById(id int) error {
	_, err := r.db.Exec("DELETE FROM Communities WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
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
