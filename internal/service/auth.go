package service

import (
	"fmt"
	"github.com/LukaGiorgadze/gonull/v2"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/model"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
}

type authService struct {
	userRepo  repository.UserRepository
	validator *validator.Validate
}

func NewAuthService(userRepo repository.UserRepository, validator *validator.Validate) AuthService {
	return &authService{
		userRepo:  userRepo,
		validator: validator,
	}
}

func (s *authService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("%w: %s", ErrValidation, "passwords don't match")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username:    req.Username,
		Email:       req.Email,
		Description: gonull.NewNullable[string](req.Description),
		Password:    string(hashed),
	}

	registered, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	var resp dto.AuthResponse

	resp.User = registered
	resp.Token = "Placeholder"

	return &resp, nil
}

func (s *authService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	err := s.validator.Var(req.Username, "email")

	var user *model.User

	if err != nil {
		user, err = s.userRepo.FindByUsername(req.Username)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.userRepo.FindByEmail(req.Username)
		if err != nil {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	var resp dto.AuthResponse

	resp.User = user
	resp.Token = "Placeholder"

	return &resp, nil
}
