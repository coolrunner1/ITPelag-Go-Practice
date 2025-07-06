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

type PostService interface {
	GetAll(start, limit int) (*dto.PostSearchResponse, error)
	GetById(id int) (*model.Post, error)
	Create(req *dto.PostCreateRequest, userId, communityId int) (*model.Post, error)
	GetAllByCommunityId(start, limit, communityId int) (*dto.PostSearchResponse, error)
	DeleteById(id, userId int) error
}

type postService struct {
	postRepo  repository.PostRepository
	validator *validator.Validate
}

func NewPostService(postRepo repository.PostRepository, validator *validator.Validate) PostService {
	return &postService{
		postRepo:  postRepo,
		validator: validator,
	}
}

func (s *postService) GetAll(start, limit int) (*dto.PostSearchResponse, error) {
	resp, err := s.postRepo.GetAll(start, limit)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrNotFound)
		}
		return nil, err
	}

	return resp, nil
}

func (s *postService) GetAllByCommunityId(start, limit, communityId int) (*dto.PostSearchResponse, error) {
	resp, err := s.postRepo.GetAllByCommunityId(start, limit, communityId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrNotFound)
		}
		return nil, err
	}

	return resp, nil
}

func (s *postService) GetById(id int) (*model.Post, error) {
	resp, err := s.postRepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w`, ErrNotFound)
		}

		return nil, err
	}

	return resp, nil
}

func (s *postService) Create(req *dto.PostCreateRequest, userId, communityId int) (*model.Post, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	res, err := s.postRepo.Create(req, userId, communityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w`, ErrNotFound)
		}
		return nil, err
	}

	return res, nil
}

func (s *postService) DeleteById(id, userId int) error {
	post, err := s.postRepo.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf(`%w`, ErrNotFound)
		}
		return err
	}

	if post.AuthorId != userId {
		return fmt.Errorf(`%w`, ErrNotAuthorized)
	}

	if err := s.postRepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}
