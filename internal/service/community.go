package service

import (
	"database/sql"
	"fmt"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator/v10"
)

type CommunityService interface {
	GetAll(start, limit int) (*dto.CommunitySearchResponse, error)
	GetById(id int) (*model.Community, error)
	Create(req *dto.CommunityCreateRequest, userId int) (*model.Community, error)
	Update(req *dto.CommunityUpdateRequest, communityId int, userId int) (*model.Community, error)
	DeleteById(id int, userId int) error
}

type communityService struct {
	communityRepo repository.CommunityRepository
	validator     *validator.Validate
}

func NewCommunityService(communityRepo repository.CommunityRepository, validator *validator.Validate) CommunityService {
	return &communityService{
		communityRepo: communityRepo,
		validator:     validator,
	}
}

func (s *communityService) GetAll(start, limit int) (*dto.CommunitySearchResponse, error) {
	resp, err := s.communityRepo.GetAll(start, limit)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *communityService) GetById(id int) (*model.Community, error) {
	resp, err := s.communityRepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w`, ErrNotFound)
		}

		return nil, err
	}

	return resp, nil
}

func (s *communityService) Create(req *dto.CommunityCreateRequest, userId int) (*model.Community, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	res, err := s.communityRepo.Create(req, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w: category doesn't exist'`, ErrNotFound)
		}
		return nil, err
	}

	return res, nil
}

func (s *communityService) Update(req *dto.CommunityUpdateRequest, communityId int, userId int) (*model.Community, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	community, err := s.communityRepo.GetById(communityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w`, ErrNotFound)
		}
	}

	if community.OwnerID != userId {
		return nil, fmt.Errorf(`%w`, ErrNotAuthorized)
	}

	res, err := s.communityRepo.Update(req, userId, communityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w'`, ErrNotFound)
		}
		return nil, err
	}

	return res, nil
}

func (s *communityService) DeleteById(id int, userId int) error {
	community, err := s.communityRepo.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf(`%w`, ErrNotFound)
		}
		return err
	}

	if community.OwnerID != userId {
		return fmt.Errorf(`%w`, ErrNotAuthorized)
	}

	if err := s.communityRepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}
