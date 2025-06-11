package service

import (
	"github.com/LukaGiorgadze/gonull/v2"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(start, limit int) ([]model.User, error)
	GetById(id int) (*model.User, error)
	Update(userRequest dto.UserUpdateRequest, id int) (*model.User, error)
	DeleteById(id int) error
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

func (s *userService) GetAll(start, limit int) ([]model.User, error) {
	return s.userRepo.GetAll(start, limit)
}

func (s *userService) GetById(id int) (*model.User, error) {
	return s.userRepo.GetById(id)
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

	if _, err = s.userRepo.GetById(id); err != nil {
		return nil, err
	}

	res, err := s.userRepo.Update(*user, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteById(id int) error {
	if _, err := s.userRepo.GetById(id); err != nil {
		return err
	}

	if err := s.userRepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}
