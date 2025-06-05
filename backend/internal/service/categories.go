package service

import (
	"context"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/labstack/gommon/log"
)

type CategoriesService interface {
	Create(ctx context.Context, req dto.CategoryCreate) error
}

type categoriesService struct {
	repo repository.CategoryRepository
}

func NewCategoriesService(repo repository.CategoryRepository) CategoriesService {
	return &categoriesService{
		repo: repo,
	}
}

func (s *categoriesService) Create(ctx context.Context, req dto.CategoryCreate) error {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(ctx, &category)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating category: %v", err))
		return err
	}

	return nil
}
