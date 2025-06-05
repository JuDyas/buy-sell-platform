package repository

import (
	"context"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
}

type categoryRepository struct {
	coll *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database, collection string) CategoryRepository {
	return &categoryRepository{
		coll: db.Collection(collection),
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	category.ID = primitive.NewObjectID()
	category.CreatedAt = time.Now()
	category.UpdatedAt = category.CreatedAt
	_, err := r.coll.InsertOne(ctx, category)

	return err
}
