package service

import (
	"context"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdvertService interface {
	Create(ctx context.Context, authorID primitive.ObjectID, req dto.AdvertCreate) (*models.Advert, error)
	Update(ctx context.Context, advertID primitive.ObjectID, req dto.AdvertUpdate) error
	GetByID(ctx context.Context, advertID primitive.ObjectID) (*models.Advert, error)
	SoftDelete(ctx context.Context, advertID primitive.ObjectID) error
}

type advertService struct {
	repo repository.AdvertRepository
}

func NewAdvertService(repo repository.AdvertRepository) AdvertService {
	return &advertService{
		repo: repo,
	}
}

func (s *advertService) Create(ctx context.Context, authorID primitive.ObjectID, req dto.AdvertCreate) (*models.Advert, error) {
	advert := &models.Advert{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Images:      req.Images,
		AuthorID:    authorID,
	}

	adv, err := s.repo.Create(ctx, advert)
	if err != nil {
		return nil, fmt.Errorf("failed to create advert: %w", err)
	}

	return adv, nil
}

func (s *advertService) Update(ctx context.Context, advertID primitive.ObjectID, req dto.AdvertUpdate) error {
	update, err := structToBsonMap(req)
	if err != nil {
		return fmt.Errorf("failed to convert struct to bson map: %w", err)
	}

	err = s.repo.Update(ctx, advertID, update)
	if err != nil {
		return fmt.Errorf("failed to update advert: %w", err)
	}

	return nil
}

func (s *advertService) GetByID(ctx context.Context, advertID primitive.ObjectID) (*models.Advert, error) {
	advert, err := s.repo.FindByID(ctx, advertID)
	if err != nil {
		return nil, fmt.Errorf("failed to get advert: %w", err)
	}

	return advert, nil
}

func (s *advertService) SoftDelete(ctx context.Context, advertID primitive.ObjectID) error {
	err := s.repo.SoftDelete(ctx, advertID)
	if err != nil {
		return fmt.Errorf("failed to delete advert: %w", err)
	}

	return nil
}
