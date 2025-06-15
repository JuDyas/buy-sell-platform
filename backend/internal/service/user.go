package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/auth"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
)

type UserService interface {
	Register(ctx context.Context, jwtSecret []byte, req dto.UserRegister) (string, error)
	Login(ctx context.Context, jwtSecret []byte, req dto.UserLogin) (string, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*dto.UserPublic, error)
	UpdateByID(ctx context.Context, id primitive.ObjectID, req dto.UserUpdate) error
	UpdateAvatar(ctx context.Context, id string, avatarURL string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetByID(ctx context.Context, id primitive.ObjectID) (*dto.UserPublic, error) {
	user, err := us.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	pubUser := &dto.UserPublic{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     user.Phone,
		Location:  user.Location,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
	}

	return pubUser, nil
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

func (us *userService) Login(ctx context.Context, jwtSecret []byte, req dto.UserLogin) (string, error) {
	user, err := us.repo.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := auth.GenerateJWT(jwtSecret, user.ID.Hex(), user.Username, int(user.Role))
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt: %w", err)
	}

	return token, nil
}

func (us *userService) UpdateAvatar(ctx context.Context, id string, avatarURL string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id to object id: %w", err)
	}

	update := bson.M{"avatar_url": avatarURL}
	err = us.repo.UpdateByID(ctx, objID, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (us *userService) UpdateByID(ctx context.Context, id primitive.ObjectID, req dto.UserUpdate) error {
	update, err := structToBsonMap(req)
	if err != nil {
		return fmt.Errorf("failed to convert struct to bson map: %w", err)
	}

	if len(update) == 0 {
		return fmt.Errorf("empty update")
	}

	err = us.repo.UpdateByID(ctx, id, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func structToBsonMap(s interface{}) (bson.M, error) {
	tmp, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(tmp, &m); err != nil {
		return nil, err
	}

	if userUpdate, ok := s.(dto.UserUpdate); ok {
		if userUpdate.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.DefaultCost)
			if err != nil {
				return nil, fmt.Errorf("failed to hash password: %w", err)
			}

			m["password_hash"] = string(hash)
		}

		delete(m, "password")
	}

	for k, v := range m {
		if str, ok := v.(string); ok && str == "" {
			delete(m, k)
		}
	}

	return m, nil
}
