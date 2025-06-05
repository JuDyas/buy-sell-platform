package service

import (
	"context"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoriesService interface {
	Create(ctx context.Context, req dto.CategoryCreate) error
	Update(ctx context.Context, id primitive.ObjectID, req dto.CategoryUpdate) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
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

func (s *categoriesService) Update(ctx context.Context, id primitive.ObjectID, req dto.CategoryUpdate) error {
	update, err := structToBsonMap(req)
	if err != nil {
		log.Error(fmt.Sprintf("Error converting struct to bson map: %v", err))
		return err
	}

	err = s.repo.Update(ctx, id, update)
	if err != nil {
		log.Error(fmt.Sprintf("Error updating category: %v", err))
		return err
	}

	return nil
}

func (s *categoriesService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting category: %v", err))
		return nil, err
	}

	return category, nil
}

func (s *categoriesService) GetAll(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting categories: %v", err))
		return nil, err
	}

	return categories, nil
}

func (s *categoriesService) Delete(ctx context.Context, id primitive.ObjectID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Error(fmt.Sprintf("Error deleting category: %v", err))
		return err
	}

	return nil
}
