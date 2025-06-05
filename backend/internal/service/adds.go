package service

import (
	"context"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var (
	fileDir = "./static/images/adverts/"
)

type AdvertService interface {
	Create(ctx context.Context, authorID primitive.ObjectID, req dto.AdvertCreate) (*models.Advert, error)
	Update(ctx context.Context, advertID primitive.ObjectID, req dto.AdvertUpdate) error
	GetByID(ctx context.Context, advertID primitive.ObjectID) (*models.Advert, error)
	SoftDelete(ctx context.Context, advertID primitive.ObjectID) error
	UploadImages([]*multipart.FileHeader) ([]string, error)
	GetAll(ctx context.Context) ([]models.Advert, error)
	GetByCategory(ctx context.Context, categoryID primitive.ObjectID) ([]models.Advert, error)
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

func (s *advertService) UploadImages(files []*multipart.FileHeader) ([]string, error) {
	var urls []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			log.Error(err)
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer src.Close()

		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Error(err)
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}

		fileName := primitive.NewObjectID().Hex() + filepath.Ext(file.Filename)
		dstPath := filepath.Join(fileDir, fileName)

		dst, err := os.Create(dstPath)
		if err != nil {
			log.Error(err)
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			log.Error(err)
			return nil, fmt.Errorf("failed to copy file: %w", err)
		}

		urls = append(urls, fileDir+fileName)
	}

	return urls, nil
}

func (s *advertService) GetAll(ctx context.Context) ([]models.Advert, error) {
	adverts, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting adverts: %v", err))
		return nil, err
	}

	return adverts, nil
}

func (s *advertService) GetByCategory(ctx context.Context, categoryID primitive.ObjectID) ([]models.Advert, error) {
	adverts, err := s.repo.GetByCategory(ctx, categoryID)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting adverts: %v", err))
		return nil, err
	}

	return adverts, nil
}
