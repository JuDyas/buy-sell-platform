package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/auth"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserNotFound          = errors.New("user not found")
)

type UserService interface {
	Register(ctx context.Context, jwtSecret []byte, req dto.UserRegister) (string, error)
	Login(ctx context.Context, req dto.UserLogin) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) Register(ctx context.Context, jwtSecret []byte, req dto.UserRegister) (string, error) {
	exist, _ := us.repo.FindByEmail(ctx, req.Email)
	if exist != nil {
		return "", ErrEmailAlreadyExists
	}

	exist, _ = us.repo.FindByUsername(ctx, req.Username)
	if exist != nil {
		return "", ErrUsernameAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         models.RoleUser,
	}

	err = us.repo.Create(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := auth.GenerateJWT(jwtSecret, user.ID.Hex(), user.Username, int(user.Role))
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt: %w", err)
	}

	return token, nil
}

func (us *userService) Login(ctx context.Context, req dto.UserLogin) (string, error) {
	return "", nil
}
