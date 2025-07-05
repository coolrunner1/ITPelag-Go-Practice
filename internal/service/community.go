package service

import (
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-playground/validator/v10"
)

type CommunityService interface {
	GetAll(start, limit int) ([]model.Community, error)
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

func (s *communityService) GetAll(start, limit int) ([]model.Community, error) {
	return s.communityRepo.GetAll(start, limit)
}
