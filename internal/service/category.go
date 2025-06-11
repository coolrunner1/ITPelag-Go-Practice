package service

import (
	"fmt"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (*model.Category, error)
	Create(req dto.CategoryRequest) (*model.Category, error)
	Update(req dto.CategoryRequest, id int) (*model.Category, error)
	DeleteById(id int) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
	validator    *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, validator *validator.Validate) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		validator:    validator,
	}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *categoryService) GetById(id int) (*model.Category, error) {
	return s.categoryRepo.GetById(id)
}

func (s *categoryService) Create(req dto.CategoryRequest) (*model.Category, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	category := model.Category{
		Title: req.Title,
	}

	created, err := s.categoryRepo.Create(category)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *categoryService) Update(req dto.CategoryRequest, id int) (*model.Category, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	category := model.Category{
		Title: req.Title,
	}

	if _, err := s.categoryRepo.GetById(id); err != nil {
		return nil, err
	}

	updated, err := s.categoryRepo.Update(category, id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *categoryService) DeleteById(id int) error {
	if _, err := s.categoryRepo.GetById(id); err != nil {
		return err
	}

	if err := s.categoryRepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}
