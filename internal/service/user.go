package service

import (
	"database/sql"
	"fmt"
	"github.com/LukaGiorgadze/gonull/v2"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(start, limit int) (*dto.UserSearchResponse, error)
	GetById(id int) (*model.User, error)
	Update(userRequest dto.UserUpdateRequest, id int) (*model.User, error)
	DeleteById(id int) error
	RestoreById(id int) (*model.User, error)
}

type userService struct {
	userRepo  repository.UserRepository
	validator *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, validator *validator.Validate) UserService {
	return &userService{
		userRepo:  userRepo,
		validator: validator,
	}
}

func (s *userService) GetAll(start, limit int) (*dto.UserSearchResponse, error) {
	return s.userRepo.GetAll(start, limit)
}

func (s *userService) GetById(id int) (*model.User, error) {
	user, err := s.userRepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`%w`, ErrNotFound)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(req dto.UserUpdateRequest, id int) (*model.User, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:    req.Username,
		Password:    string(hashed),
		Email:       req.Email,
		Description: gonull.NewNullable[string](req.Description),
	}

	res, err := s.userRepo.Update(*user, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrNotFound)
		}

		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteById(id int) error {
	user, err := s.userRepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%w", ErrNotFound)
		}

		return err
	}

	if user.DeletedAt.Valid {
		return fmt.Errorf("%w", ErrAlreadyDeleted)
	}

	if err := s.userRepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}

func (s *userService) RestoreById(id int) (*model.User, error) {
	user, err := s.userRepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrNotFound)
		}

		return nil, err
	}

	if !user.DeletedAt.Valid {
		return nil, fmt.Errorf("%w", ErrNotDeleted)
	}

	user, err = s.userRepo.RestoreById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
