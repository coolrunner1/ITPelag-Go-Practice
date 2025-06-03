package service

import (
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/repository"
)

type UserService interface {
	GetAll(start, limit int) ([]model.User, error)
	GetById(id int) (*model.User, error)
	//Create(c model.Category) (*model.Category, error)
	//Update(c model.Category, id int) (*model.Category, error)
	//DeleteById(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAll(start, limit int) ([]model.User, error) {
	return s.userRepo.GetAll(start, limit)
}

func (s *userService) GetById(id int) (*model.User, error) {
	return s.userRepo.GetById(id)
}
